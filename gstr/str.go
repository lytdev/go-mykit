package gstr

import (
	"errors"
	"regexp"
	"strings"
	"unsafe"
)

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

// IsEmpty 是否为空
func IsEmpty(val string) bool {
	s := strings.TrimSpace(val)
	return len(s) == 0
}

// ReplaceAll 将给定的所有字符统一替换成指定的字符
func ReplaceAll(old string, target string, srcTag ...string) (string, error) {
	var builder strings.Builder
	builder.WriteString("[")
	for _, tag := range srcTag {
		builder.WriteString(tag)
	}
	builder.WriteString("]")
	return ReplaceStringByRegex(old, builder.String(), target)
}

// ReplaceStringByRegex 通过正则表达式把字符串替换掉
func ReplaceStringByRegex(src, rule, target string) (string, error) {
	reg, err := regexp.Compile(rule)
	if reg == nil || err != nil {
		return "", errors.New("正则表达式编译错误:" + err.Error())
	}
	return reg.ReplaceAllString(src, target), nil
}

// SplitIgnoreBlank 切分字符串,忽略空内容
func SplitIgnoreBlank(s, sep string) []string {
	if s == "" || len(s) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	split := strings.Split(s, sep)
	for _, str := range split {
		if str != "" && len(str) > 0 {
			result = append(result, str)
		}
	}
	return result
}
