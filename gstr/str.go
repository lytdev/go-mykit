package gstr

import (
	"crypto/rand"
	"strings"
	"unsafe"
)

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

const (
	LettersLetter          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LettersUpperCaseLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LettersNumber          = "0123456789"
	LettersNumberNoZero    = "23456789"
	LettersSymbol          = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
)

// RandString 随机字符串
func RandString(n int, letters ...string) (string, error) {

	lettersDefaultValue := LettersLetter + LettersNumber + LettersSymbol

	if len(letters) > 0 {
		lettersDefaultValue = ""
		for _, letter := range letters {
			lettersDefaultValue = lettersDefaultValue + letter
		}
	}

	bytes := make([]byte, n)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = lettersDefaultValue[b%byte(len(lettersDefaultValue))]
	}

	return string(bytes), nil
}
