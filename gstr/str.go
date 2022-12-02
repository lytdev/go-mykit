package gstr

import (
	"errors"
	"regexp"
	"strings"
)

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

// SubStr 截取字符串的长度
// @param str 原字符串
// @param start 起始下标,负数从尾部开始,-1为最后一个
// @param length 截取长度,负数表示截取到末尾
func SubStrByLen(str string, start int, length int) (result string) {
	s := []rune(str)
	total := len(s)
	if total == 0 {
		return
	}
	// 允许从尾部开始计算
	if start < 0 {
		start = total + start
		if start < 0 {
			return
		}
	}
	if start > total {
		return
	}
	// 到末尾
	if length < 0 {
		length = total
	}
	end := start + length
	if end > total {
		result = string(s[start:])
	} else {
		result = string(s[start:end])
	}
	return
}

// SubStrByStr
// @param str 原字符串
// @param indexStr 切分的字符
// @param fol 切分字符的第一次出现还是最后一次出现,1:第一次;2:最后一次;
// @param fob 向前还是向后截取,1:向前;2:向后;
func SubStrByStr(str, indexStr string, fol, fob int) (result string) {
	//默认第一次
	index := strings.Index(str, indexStr)
	if fol == 2 {
		index = strings.LastIndex(str, indexStr)
	}
	result = str[index:]
	if fob == 1 {
		result = str[0:index]
	}
	return
}
