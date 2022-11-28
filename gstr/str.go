package gstr

import "unsafe"

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
