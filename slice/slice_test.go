package slice

import (
	"fmt"
	"testing"
	"time"
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

func TestMergeInt(t *testing.T) {
	list1 := []int{1,2,3}
	list4 := []int{5,6,7}
	for i:=0;i<500000;i++{
		list1=append(list1, i)
	}
	for i:=0;i<1000000;i++{
		list4=append(list4, i)
	}
	t1 := time.Now().UnixNano()
	list5 := MergeInt(list1, list4)
	_ =list5
	t2 := time.Now().UnixNano()
	fmt.Println(t2-t1, t1, t2)
}

func TestHMergeInt(t *testing.T) {
	list1 := []int{1,2,3}
	list4 := []int{5,6,7}
	for i:=0;i<10000000;i++{
		list1=append(list1, i)
	}
	for i:=0;i<1000000;i++{
		list4=append(list4, i)
	}
	t1 := time.Now().UnixNano()
	//list5 := HMergeInt(list4, list1) //大 小
	list5 := HMergeInt(list1, list4) //小 大
	_ =list5
	t2 := time.Now().UnixNano()
	fmt.Println(time.Duration(t2-t1), t1, t2)
}
