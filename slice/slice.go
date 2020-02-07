package slice

import (
	"reflect"
	"strconv"
	"strings"
)

//判断切片中是否存在v int int64 string
func Contain(sl interface{}, v interface{}) bool {
	sType := reflect.TypeOf(sl)
	sValue := reflect.ValueOf(sl)
	vValue := reflect.ValueOf(v)
	if sType.Kind() != reflect.Slice {
		return false
	}
	if sValue.Index(0).Kind() != vValue.Kind() {
		return false
	}

	switch sl.(type) {
	case []int:
		vv := v.(int)
		vsl := sl.([]int)
		for _, vs := range vsl {
			if vv == vs {
				return true
			}
		}
	case []int64:
		vv := v.(int64)
		vsl := sl.([]int64)
		for _, vs := range vsl {
			if vv == vs {
				return true
			}
		}

	case []string:
		vv := v.(string)
		vsl := sl.([]string)
		for _, vs := range vsl {
			if vv == vs {
				return true
			}
		}
	default:
		return false
	}
	return false
}

func ContainInt(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func ContainInt64(sl []int64, v int64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func ContainString(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// SliceMerge merges interface slices to one slice.
func Merge(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

func MergeInt(slice1, slice2 []int) (c []int) {
	c = append(slice1, slice2...)
	return
}

//无序高性能
func HMergeInt(slice1, slice2 []int) []int {
	if len(slice1) < len(slice2) {
		return append(slice2, slice1...)
	}
	return append(slice1, slice2...)
}

func MergeInt64(slice1, slice2 []int64) []int64 {
	slice1 = append(slice1, slice2...)
	return slice1
}

func HMergeInt64(slice1, slice2 []int64) []int64 {
	if len(slice1) < len(slice2) {
		return append(slice2, slice1...)
	}
	return append(slice1, slice2...)
}

func MergeString(slice1, slice2 []string) []string {
	return append(slice1, slice2...)
}

func HMergeString(slice1, slice2 []string) []string {
	if len(slice1) < len(slice2) {
		return append(slice2, slice1...)
	}
	return append(slice1, slice2...)
}

func UniqueInt(s []int) []int {
	size := len(s)
	if size == 0 {
		return []int{}
	}

	ret := make([]int, 0, size)
	m := make(map[int]bool, size)
	for _, v := range s {
		if _, ok := m[v]; !ok{
			ret = append(ret, v)
			m[v] = true
		}
	}
	return ret
}

func UniqueInt64(s []int64) []int64 {
	size := len(s)
	if size == 0 {
		return []int64{}
	}

	ret := make([]int64, 0, size)
	m := make(map[int64]bool, size)
	for _, v := range s {
		if _, ok := m[v]; !ok{
			ret = append(ret, v)
			m[v] = true
		}
	}
	return ret
}

func UniqueString(s []string) []string {
	size := len(s)
	if size == 0 {
		return []string{}
	}

	ret := make([]string, 0, size)
	m := make(map[string]bool, size)
	for _, v := range s {
		if _, ok := m[v]; !ok{
			ret = append(ret, v)
			m[v] = true
		}
	}
	return ret
}

func SumInt64(sl []int64) (sum int64) {
	for _, v := range sl {
		sum += v
	}
	return
}

func SumInt(sl []int) (sum int) {
	for _, v := range sl {
		sum += v
	}
	return
}

func SumFloat64(sl []float64) (sum float64) {
	for _, v := range sl {
		sum += v
	}
	return
}

//int 转 string
func IntToInString(list []int) string {
	s := make([]string, 0, len(list))

	for _,v := range list {
		s = append(s, strconv.Itoa(v))
	}
	return strings.Join(s, ",")
}

//int64 change string
func Int64ToInString(list []int64) string {
	s := make([]string, 0, len(list))
	for _, o := range list {
		s = append(s, strconv.FormatInt(o, 10))
	}
	return strings.Join(s, ",")
}

