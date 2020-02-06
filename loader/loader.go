package loader

import (
	"compress/gzip"
	"container/list"
	"crypto/tls"
	// "encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
	"bytes"
	"golang.org/x/net/publicsuffix"
	"net/http/cookiejar"
)

var (
	mobileUserAgentS = []string{
		"Mozilla/5.0 (Linux; U; Android 4.0.2; en-us; Galaxy Nexus Build/ICL53F) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko)",
		"Mozilla/5.0 (iPad; U; CPU OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko)",
		"Mozilla/5.0 (Linux; U; Android 2.3.5; zh-cn; MI-ONE Plus Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 2.3.3; zh-cn; HTC_WildfireS_A510e Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	}
	pcUserAgentS = []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:21.0) Gecko/20130331 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.2; WOW64; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.11 (KHTML, like Gecko) Ubuntu/11.10 Chromium/27.0.1453.93 Chrome/27.0.1453.93 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.94 Safari/537.36",
	}

	androidUserAgentS = []string{
		"Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 2.2; en-gb; GT-P1000 Build/FROYO) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 4.0.2; en-us; Galaxy Nexus Build/ICL53F) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 2.3.5; zh-cn; MI-ONE Plus Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 2.3.3; zh-cn; HTC_WildfireS_A510e Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M032 Build/IML74K) AppleWebKit/533.1 (KHTML, like Gecko)Version/4.0 MQQBrowser/4.1 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M032 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M032 Build/IML74K) UC AppleWebKit/534.31 (KHTML, like Gecko) Mobile Safari/534.31",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M032 Build/IML74K) UC AppleWebKit/534.31 (KHTML, like Gecko) Mobile Safari/534.31",
	}
)

var (
	proxyList *list.List = list.New()
	proxyListLock bool       = false
)

type Loader struct {
	Runing    int
	proxyURL  string
	transport *http.Transport
	myHeader  map[string]string
}

func NewLoader() *Loader {
	loader := &Loader{
		myHeader: map[string]string{
			"Accept-Charset":  "utf-8",
			"Accept-Encoding": "gzip, deflate, sdch",
			"Content-Type":    "application/x-www-form-urlencoded",
			"Connection":      "close",
		},
		proxyURL:"",
	}
	loader.transport = loader.getTransport()
	loader.WithMobileAgent()
	return loader
}

func (loader *Loader) SetHeader(head, value string) *Loader {
	loader.myHeader[head] = value
	return loader
}

func (loader *Loader) WithMobileAgent() *Loader {
	num := RandInt(0, len(mobileUserAgentS)-1)
	loader.myHeader["User-Agent"] = mobileUserAgentS[num]
	return loader
}

func (loader *Loader) WithAndroidAgent() *Loader {
	num := RandInt(0, len(androidUserAgentS)-1)
	loader.myHeader["User-Agent"] = androidUserAgentS[num]
	return loader
}

func (loader *Loader) WithPcAgent() *Loader {
	num := RandInt(0, len(pcUserAgentS)-1)
	loader.myHeader["User-Agent"] = pcUserAgentS[num]
	return loader
}

func (loader *Loader) WithHttpProxy(proxy string) *Loader {
	loader.proxyURL = fmt.Sprintf("http://%s", proxy)
	proxyUrl, _ := url.Parse(loader.proxyURL)
	loader.transport.Proxy = http.ProxyURL(proxyUrl)
	return loader
}

func (loader *Loader) getRequest(target, method string, data interface{}) *http.Request {
	var request *http.Request
	if strings.ToUpper(method) == "POST" {
		switch data.(type) {
	    case url.Values:
	    	encodeData := (data.(url.Values)).Encode()
	    	request, _ = http.NewRequest("POST", target, strings.NewReader(encodeData))	  
	    	// request.Header.Set("Content-Length", string(len(encodeData)))  	
	    case string:
	    	encodeData := data.(string)
	    	request, _ = http.NewRequest("POST", target, strings.NewReader(encodeData))
	    	// request.Header.Set("Content-Length", string(len(encodeData)))
	    case []byte:
	    	encodeData := data.([]byte)
	    	request, _ = http.NewRequest("POST", target, bytes.NewReader(encodeData))
	    	// request.Header.Set("Content-Length", string(len(encodeData)))
	    default:
	    }
	} else {
		request, _ = http.NewRequest(method, target, nil)
	}
	request.Close = true

	for h, v := range loader.myHeader {
		request.Header.Set(h, v)
	}
	return request
}

func (loader *Loader) getTransport() *http.Transport {
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, time.Second*30)
		},
		TLSClientConfig: &tls.Config{
			MinVersion:         tls.VersionTLS10,
			InsecureSkipVerify: true,
		},
		DisableKeepAlives: true,
	}
	return transport
}

func (loader *Loader) SetTransport(transport *http.Transport) *Loader {
	loader.transport = transport
	return loader
}

func (loader *Loader) Get(target string) (*http.Response, []byte, error) {
	return loader.Send(target, "GET", nil)
}

func (loader *Loader) Post(target string, data interface{}) (*http.Response, []byte, error) {
	return loader.Send(target, "POST", data)
}

func (loader *Loader) Send(target, method string, data interface{}) (*http.Response, []byte, error) {
	loader.Runing++

	options := cookiejar.Options{
        PublicSuffixList: publicsuffix.List,
    }
    jar, err := cookiejar.New(&options)
    if err != nil {
        loger.D("Loader.Send", err.Error())
    }

	client := &http.Client{
		Transport: loader.transport,
		Jar: jar,
	}

	request := loader.getRequest(target, method, data)
	resp, err := client.Do(request)
	if err != nil {
		loader.Runing--
		return nil, nil, err
	}

	defer resp.Body.Close()
	loger.D("[Loader.Send][", resp.StatusCode, "] Loader [", target, "] with proxy", loader.proxyURL)

	if resp.StatusCode != http.StatusOK {
		loader.Runing--
		return nil, nil, err
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			loader.Runing--
			return nil, nil, err
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := ioutil.ReadAll(reader)

	if err != nil {
		loader.Runing--
		return nil, nil, err
	}
	loader.Runing--
	return resp, body, nil
}
