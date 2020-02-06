package conv

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
)

//byte 转 int64 大字节序
func BytesToInt64Big(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

//byte 转 int64 小字节序
func BytesToInt64Little(buf []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buf))
}

//string to int
func StringToInt(value string) (i int) {
	i, _ = strconv.Atoi(value)
	return
}

func StringToInt64(value string) (i int64) {
	i, _ = strconv.ParseInt(value, 10, 64)
	return
}

// convert any numeric value to int64
// 任意类型转int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	case float32,float64:
		d = int64(val.Uint())

	case string:
		d, err = strconv.ParseInt(val.String(), 10, 64)
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}
