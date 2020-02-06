package conv

import (
	"fmt"
	"testing"
)

func TestInt64ToBytesBig(t *testing.T) {
	var i int64
	i = 10001
	rs := Int64ToBytesBig(i)
	fmt.Println(rs)
}

func TestInt64ToBytesLittle(t *testing.T) {
	var i int64
	i = 10001
	rs := Int64ToBytesLittle(i)
	fmt.Println(rs)
}

