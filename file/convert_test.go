package file

import (
	"testing"
	"fmt"
)

func TestToUint64(t *testing.T) {
	n,err:=ToUint64("aaa.txt")
	fmt.Println(n, err)
}
