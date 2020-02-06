package conv

import (
	"fmt"
	"testing"
)

func TestBytesToInt64Big(t *testing.T) {
	b := []byte{1,2,3,4,5,7,8,9,10}
	i64 := BytesToInt64Big(b)
	fmt.Println(i64)
}

func TestBytesToInt64Little(t *testing.T) {
	b := []byte{1,2,3,4,5,7,8,9,10}
	i64 := BytesToInt64Little(b)
	fmt.Println(i64)
}
