package array

import (
	"strconv"
	"strings"
)

//int 转 string
func IntArrToInString(list []int) string {
	s := make([]string, 0, len(list))

	for _,v := range list {
		s = append(s, strconv.Itoa(v))
	}
	return strings.Join(s, ",")
}

//int64 change string
func Int64ArrToInString(list []int64) string {
	s := make([]string, 0, len(list))
	for _, o := range list {
		s = append(s, strconv.FormatInt(o, 10))
	}
	return strings.Join(s, ",")
}

//string arr change int
func StringArrToInString(s []string) string {
	return `"` + strings.Join(s, `","`) + `"`
}

//判断元素是否包含
func InArray(s string, arr []string) bool {
	for _, val := range arr {
		if s == val {
			return true
		}
	}
	return false
}
