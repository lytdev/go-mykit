package gconv

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// ToString 将接口数据类型转换为string
func ToString(i interface{}) string {
	v, _ := ToStringE(i)
	return v
}

// https://pkg.go.dev/strings
//string与[]byte的直接转换是通过底层数据copy实现的,存在性能损耗,不过一般不影响

// StrToByteArr 字符串转字节数组
func StrToByteArr(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// ByteArrToStr 字节数组转字符串
func ByteArrToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ArrToStr 将数组转化为字符串
func ArrToStr(array interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

// StrToFloat64 字符串转化为浮点类型,支持指定精度
func StrToFloat64(str string, len int) (float64, error) {
	lenStr := "%." + strconv.Itoa(len) + "f"
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(fmt.Sprintf(lenStr, value), 64)
}
