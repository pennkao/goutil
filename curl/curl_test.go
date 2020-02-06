package curl

import (
	"fmt"
	"testing"
)

func TestHttpGet(t *testing.T) {
	url := "http://www.baidu.com"
	resp, err := HttpGet(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
