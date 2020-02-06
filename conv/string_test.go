package conv
import (
	"fmt"
	"testing"
)
func TestToString(t *testing.T){
	type Msg struct {
		Name string
		Age int
	}
	//m := Msg{"dafad", 10}
	mp := map[string]interface{}{"adf":222, "dafasdf":"adfasdf"}
	v := ToString(mp)
	fmt.Println(v)
}
