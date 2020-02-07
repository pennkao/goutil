package slice

import (
	"fmt"
	"testing"
)
type B string
func TestContain(t *testing.T) {
	s := []int8{1,2,3,4}
	v := int8(1)
	r := Contain(s, v)
	fmt.Println(r)
}
func TestSliceContainsInt(t *testing.T) {
	s := []int{1,2,3,4}
	v := 1
	r := ContainInt(s, v)
	fmt.Println(r)
}

func TestUniqueInt(t *testing.T) {
	s := []int{1,2,3,4,5,1,2,34,12,1}
	s = UniqueInt(s)
	fmt.Println(s)
}

func TestUniqueInt64(t *testing.T) {
	s := []int64{1,2,3,4,5,1,2,34,12,1}
	s = UniqueInt64(s)
	fmt.Println(s)
}

func TestUniqueString(t *testing.T) {
	s := []string{"c", "a", "b", "c", "a", "b"}
	s = UniqueString(s)
	fmt.Println(s)
}
