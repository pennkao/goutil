package extstrings

import (
	"fmt"
	"testing"
)

func TestTrimRightSpace(t *testing.T) {
	str := "aaabc\n\t\r"
	rs := TrimRightSpace(str)
	fmt.Println(rs)
}

func TestRandomString(t *testing.T) {
	str := RandomString(5)
	fmt.Println(str)
}

func TestUcfirst(t *testing.T) {
	v := Ucfirst("你哈")
	fmt.Println(v)
}
func TestLcfirst(t *testing.T) {
	v := Lcfirst("xxxxx")
	fmt.Println(v)
}

func TestWordUpper(t *testing.T) {
	v := WordUpper("weight")
	fmt.Println(v)
}