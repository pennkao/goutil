package extnet

import (
	"testing"
	"fmt"
)

func TestGetPublicIp(t *testing.T) {
	p, err:= GetPublicIp()
	fmt.Println(p,err)
}

func TestLocalIPv4s(t *testing.T) {
	p, err := LocalIPv4s()
	fmt.Println(p, err)
}

func TestGetIPv4ByInterface(t *testing.T){
	p,err := GetIPv4ByInterface("en0")
	fmt.Println(p, err)
}
