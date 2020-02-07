# go 常用hash函数

## 示例

```
package main

import (
	"fmt"
	"github.com/pennkao/goutil/slice"
)

func main() {
    
    list1 := []int{1,2,3,4}
    v1 := 1
    r := slice.Contain(list1, v1)
    fmt.Println(r) //true

	s := []string{"c", "a", "b", "c", "a", "b"}
	s = slice.UniqueString(s)
	fmt.Println(s)  //[c a b]

    
}

```