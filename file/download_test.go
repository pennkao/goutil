package file

import (
	"testing"
	"fmt"
)

func TestDownload(t *testing.T) {
	v, err := Download("aaa.txt", "http://www.baidu.com", nil)
	fmt.Println(v, err)
}

func TestWget(t *testing.T) {
	Wget("bbb.txt", "http://www.baidu.com")

}