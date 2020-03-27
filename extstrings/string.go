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

func Ucfirst(str string) string {
	if len(str) < 1 {
		return str
	}
	strArr := []rune(str)
	if len(strArr) < 1 {
		return str
	}

	if strArr[0] >= 97 && strArr[0] <= 122 {
		strArr[0] -= 32
	}

	return string(strArr)
}

func Lcfirst(str string) string {
	if len(str) < 1 {
		return str
	}

	strArr := []rune(str)
	if len(strArr) < 1 {
		return str
	}

	len := len(str)
	if strArr[len-1] >= 97 && strArr[len-1] <= 122 {
		strArr[len-1] -= 32
	}

	return string(strArr)
}

func WordUpper(str string) string{
	if len(str) < 1 {
		return str
	}

	strArr := make([]string, 0, 50)
	if strings.Contains(str, "-") {
		strArr = strings.Split(str, "-")
	} else if strings.Contains(str, "_"){
		strArr = strings.Split(str, "_")
	}
	if len(strArr) < 1 {
		return str
	}
	for i, v := range strArr {
		strArr[i] = Ucfirst(v)
	}

	return strings.Join(strArr, "")
}
