package extnet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
)

// Constants to choose which version of SOCKS protocol to use.
const (
	SOCKS4  = 3
	SOCKS4A = 3
	SOCKS5  = 4
)

type SocksAuth struct {
	Username, Password string
}

const (
	socks5IP4    = 1
	socks5Domain = 3
	socks5IP6    = 4
)

var socks5Errors = []string{
	"",
	"general failure",
	"connection forbidden",
	"network unreachable",
	"host unreachable",
	"connection refused",
	"TTL expired",
	"command not supported",
	"address type not supported",
}

// DialSocksProxy returns the dial function to be used in http.Transport object.
// Argument socksType should be one of SOCKS4, SOCKS4A and SOCKS5.
// Argument proxy should be in this format "127.0.0.1:1080".
func DialSocksProxy(socksType int, proxy string, remoteUrl string, auth *SocksAuth) (net.Conn, error) {
	if socksType == SOCKS5 {
		return dialSocks5(proxy, remoteUrl, auth)
	}

	// SOCKS4, SOCKS4A
	return dialSocks4(socksType, proxy, remoteUrl)
}

func dialSocks5(proxy, remoteUrl string, auth *SocksAuth) (conn net.Conn, err error) {
	// dial TCP
	conn, err = net.Dial("tcp", proxy)
	if err != nil {
		return
	}

	host, port, err := splitHostPort(remoteUrl)
	// version identifier/method selection request
	req := make([]byte, 0, 6+len(host))
	req = append(req, 5)
	if auth != nil {
		req = append(req, 2, 0, 2)
	} else {
		req = append(req, 1, 0)
	}

	resp, err := sendReceive(conn, req)

	if err != nil {
		return nil, err
	}

	if len(resp) != 2 {
		return nil, errors.New("Server does not respond properly.")
	}
	if resp[0] != 5 {
		return nil, errors.New("Server does not support Socks 5.")
	}
	if resp[1] == 0 { // no auth
		return nil, errors.New("proxy: SOCKS5 proxy requires authentication")
	}

	if req[1] == 2 {
		buf := make([]byte, 0)
		buf = append(buf, 1)
		buf = append(buf, uint8(len(auth.Username)))
		buf = append(buf, auth.Username...)
		buf = append(buf, uint8(len(auth.Password)))
		buf = append(buf, auth.Password...)

		resp, err := sendReceive(conn, buf)

		if err != nil {
			return nil, err
		}

		if len(resp) != 2 {
			err = errors.New("failed to read authentication reply from SOCKS5 proxy at " + err.Error())
			return nil, err
		}

		if resp[1] != 0 {
			err = errors.New("SOCKS5 proxy rejected username/password ")
			return nil, err
		}
	}

	// detail request
	buf := make([]byte, 0)
	buf = append(buf, 5, 1, 0 /* reserved */)

	if ip := net.ParseIP(host); ip != nil {
		if ip4 := ip.To4(); ip4 != nil {
			buf = append(buf, socks5IP4)
			ip = ip4
		} else {
			buf = append(buf, socks5IP6)
		}
		buf = append(buf, ip...)
	} else {
		buf = append(buf, socks5Domain)
		buf = append(buf, byte(len(host)))
		buf = append(buf, host...)
	}
	buf = append(buf, byte(port>>8), byte(port))

	if _, err := conn.Write(buf); err != nil {
		return nil, errors.New("proxy: failed to write connect request to SOCKS5 proxy at " + err.Error())
	}

	if _, err := io.ReadFull(conn, buf[:4]); err != nil {
		return nil, errors.New("proxy: failed to read connect reply from SOCKS5 proxy at " + err.Error())
	}

	failure := "unknown error"
	if int(buf[1]) < len(socks5Errors) {
		failure = socks5Errors[buf[1]]
	}

	if len(failure) > 0 {
		return nil, errors.New("proxy: SOCKS5 proxy failed to connect: " + failure)
	}

	bytesToDiscard := 0
	switch buf[3] {
	case socks5IP4:
		bytesToDiscard = net.IPv4len
	case socks5IP6:
		bytesToDiscard = net.IPv6len
	case socks5Domain:
		_, err := io.ReadFull(conn, buf[:1])
		if err != nil {
			return nil, errors.New("proxy: failed to read domain length from SOCKS5 proxy at " + err.Error())
		}
		bytesToDiscard = int(buf[0])
	default:
		return nil, errors.New("proxy: got unknown address type " + strconv.Itoa(int(buf[3])) + " from SOCKS5 proxy ")
	}

	if cap(buf) < bytesToDiscard {
		buf = make([]byte, bytesToDiscard)
	} else {
		buf = buf[:bytesToDiscard]
	}
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, errors.New("proxy: failed to read address from SOCKS5 proxy " + err.Error())
	}

	// Also need to discard the port number
	if _, err := io.ReadFull(conn, buf[:2]); err != nil {
		return nil, errors.New("proxy: failed to read port from SOCKS5 proxy " + err.Error())
	}

	return conn, nil
}

func dialSocks4(socksType int, proxy string, remoteUrl string) (conn net.Conn, err error) {
	// dial TCP
	conn, err = net.Dial("tcp", proxy)
	if err != nil {
		return
	}

	// connection request
	host, port, err := splitHostPort(remoteUrl)
	if err != nil {
		return
	}
	ip := net.IPv4(0, 0, 0, 1).To4()
	if socksType == SOCKS4 {
		ip, err = lookupIP(host)
		if err != nil {
			return
		}
	}
	req := []byte{
		4,                          // version number
		1,                          // command CONNECT
		byte(port >> 8),            // higher byte of destination port
		byte(port),                 // lower byte of destination port (big endian)
		ip[0], ip[1], ip[2], ip[3], // special invalid IP address to indicate the host name is provided
		0, // user id is empty, anonymous proxy only
	}
	if socksType == SOCKS4A {
		req = append(req, []byte(host+"\x00")...)
	}

	resp, err := sendReceive(conn, req)
	if err != nil {
		return
	} else if len(resp) != 8 {
		err = errors.New("Server does not respond properly.")
	}
	switch resp[1] {
	case 90:
		// request granted
	case 91:
		err = errors.New("Socks connection request rejected or failed.")
	case 92:
		err = errors.New("Socks connection request rejected becasue SOCKS server cannot connect to identd on the client.")
	case 93:
		err = errors.New("Socks connection request rejected because the client program and identd report different user-ids.")
	default:
		err = errors.New("Socks connection request failed, unknown error.")
	}
	return
}

func sendReceive(conn net.Conn, req []byte) (resp []byte, err error) {
	_, err = conn.Write(req)
	if err != nil {
		return
	}
	resp, err = readAll(conn)
	return
}

func readAll(conn net.Conn) (resp []byte, err error) {
	resp = make([]byte, 1024)
	n, err := conn.Read(resp)
	resp = resp[:n]
	return
}

func lookupIP(host string) (ip net.IP, err error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return
	}
	if len(ips) == 0 {
		err = errors.New(fmt.Sprintf("Cannot resolve host: %s.", host))
		return
	}
	ip = ips[0].To4()
	if len(ip) != net.IPv4len {
		fmt.Println(len(ip), ip)
		err = errors.New("IPv6 is not supported by SOCKS4.")
		return
	}
	return
}

func splitHostPort(addr string) (host string, port uint16, err error) {
	host, portStr, err := net.SplitHostPort(addr)
	portInt, err := strconv.ParseUint(portStr, 10, 16)
	port = uint16(portInt)
	return
}
