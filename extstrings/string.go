package extstrings

import (
	"math/rand"
	"strings"
	"time"
)

//去掉换行,空格,回车,制表
func TrimRightSpace(s string) string {
	return strings.TrimRight(string(s), "\r\n\t ")
}


//随机字符串
func RandomString(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
