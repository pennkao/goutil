package array

import (
	"fmt"
	"testing"
	"time"
)

func TestIntArrToInString(t *testing.T) {
	t1 := time.Now().UnixNano()
	i := []int{1, 2, 3, 4, 5}
	t.Log("dfdsf")
	rs := IntArrToInString(i)
	t2 := time.Now().UnixNano()
	fmt.Println(t2-t1)
	fmt.Println(rs, time.Microsecond)
}
func TestInt64ArrToInString(t *testing.T) {
	i := []int64{11, 12, 13, 14}
	rs := Int64ArrToInString(i)
	fmt.Println(rs)
}
func TestStringArrToInString(t *testing.T) {
	i := []string{"a", "b", "c"}
	rs := StringArrToInString(i)
	fmt.Println(rs)
}
func TestInArray(t *testing.T) {
	bool := InArray("1", []string{"0", "1", "12"})
	fmt.Println(bool)
}
