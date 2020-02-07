package page

import (
	"testing"
	"fmt"
)

func TestTest(t *testing.T) {
	p := new(Page)
	p.Page = 1
	p.Size = 10

	p.SetTotal(100)
	fmt.Println(p.GetCount())
	fmt.Println(p.OffSet, p.Limit())
}