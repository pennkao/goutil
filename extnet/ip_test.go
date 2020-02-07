package extnet

import (
	"testing"
	"fmt"
)

func TestGetPulicIP(t *testing.T) {
	p := GetPulicIP()
	fmt.Println(p)
}

func TestLocalIPv4s(t *testing.T) {
	p, err := LocalIPv4s()
	fmt.Println(p, err)
}

func TestGetIPv4ByInterface(t *testing.T){
	p,err := GetIPv4ByInterface("en0")
	fmt.Println(p, err)
}
