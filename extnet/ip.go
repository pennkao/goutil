package extnet

import (
	"strconv"
	"strings"
)

func Ip2Int(ip string) int {
	lisit := strings.FieldsFunc(ip, func(c rune) bool { return c == '.' })

	ip1_str_int, _ := strconv.Atoi(lisit[0])
	ip2_str_int, _ := strconv.Atoi(lisit[1])
	ip3_str_int, _ := strconv.Atoi(lisit[2])
	ip4_str_int, _ := strconv.Atoi(lisit[3])
	return ip1_str_int<<24 | ip2_str_int<<16 | ip3_str_int<<8 | ip4_str_int
}
