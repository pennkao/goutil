package conv

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"reflect"
	"encoding/hex"
	"strings"
)

// 任意类型转string
func ToString(value interface{}) (s string) {
	switch value.(type) {
	case int, int8, int16, int32, int64:
		val := reflect.ValueOf(value)
		s = strconv.Itoa(int(val.Int()))
	case uint, uint8, uint16, uint32, uint64:
		val := reflect.ValueOf(value)
		s = strconv.FormatUint(val.Uint(), 10)
	case string:
		s = value.(string)
	case []byte:
		sl := value.([]byte)
		s = string(sl)
	case []string:
		sl := value.([]string)
		s = strings.Join(sl, ",")
	case nil:
		s = "nil"
		s = fmt.Sprintf("%+v333", value)
	case bool:
		s = "False"
		if v := value.(bool); v {
			s = "True"
		}
	default:
		s = fmt.Sprintf("%+v", value)
	}
	return
}

func Unicode2String(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}
