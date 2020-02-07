package extnet

import (
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func IpString2Int(ip string) int {

	bytes := strings.Split(ip, ".")
	ip1_str_int, _ := strconv.Atoi(bytes[0])
	ip2_str_int, _ := strconv.Atoi(bytes[1])
	ip3_str_int, _ := strconv.Atoi(bytes[2])
	ip4_str_int, _ := strconv.Atoi(bytes[3])
	return ip1_str_int<<24 | ip2_str_int<<16 | ip3_str_int<<8 | ip4_str_int
}

func IpString2Int64(ip string)int64{
	bits := strings.Split(ip, ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func IpNet2Int64(ipNet net.IP) int64 {
	return IpString2Int64(ipNet.String())
}

func IpNet2Int(ipNet net.IP) int {
	return IpString2Int(ipNet.String())
}


func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

//func GetPulicIP() net.IP {
//	conn, _ := net.DialTimeout("udp", "220.181.38.149:80", 1*time.Second)
//
//	idx := strings.Split(conn.LocalAddr().String(), ":")[0]
//	//关闭连接
//	conn.Close()
//	_ip := net.ParseIP(idx)
//	if _ip != nil {
//		return _ip
//	}
//
//	//logger.Fatal("获取外网地址失败",idx)
//	return nil
//}

//func GetPulicIP() string {
//	conn, _ := net.Dial("udp", "8.8.8.8:80")
//	defer conn.Close()
//	localAddr := conn.LocalAddr().String()
//	idx := strings.LastIndex(localAddr, ":")
//	return localAddr[0:idx]
//}

func GetPublicIp() (string,error) {
	//有点慢
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	return string(content), err
}

func LocalIPv4s() ([]string, error) {
	var ips []string
	ipList, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, ip := range ipList {
		if ipNet, ok := ip.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, nil
}

func GetIPv4ByInterface(name string) ([]string, error) {
	var ips []string

	iFace, err := net.InterfaceByName(name)
	if err != nil {
		return nil, err
	}

	ipList, err := iFace.Addrs()
	if err != nil {
		return nil, err
	}

	for _, a := range ipList {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, nil
}
