package file

import (
	"fmt"
	"testing"
)

func TestToUint64(t *testing.T) {
	n,err:=ToUint64("aaa.txt")
	fmt.Println(n, err)
}
