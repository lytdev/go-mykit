package grandom

import (
	"math/rand"
	"strings"
)

const (
	lettersLetter          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	enAlphabetLen          = len(lettersLetter)
	lettersUpperCaseLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerEnLetter          = "abcdefghijklmnopqrstuvwxyz"
	lowerEnLetterLen       = len(lowerEnLetter)
	lettersNumber          = "0123456789"
	lettersNumberNoZero    = "23456789"
	lettersSymbol          = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
)

// RandomString 随机字符串
func RandomString(n int, letters ...string) (string, error) {
	lettersDefaultValue := lettersLetter + lettersNumber + lettersSymbol
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

// RandomStr 随机字符串
func RandomStr(l uint64) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = lettersLetter[RandInt(enAlphabetLen)]
	}
	return string(b)
}

// RandomLowerStr 随机字符串小写
func RandomLowerStr(l uint32) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = lowerEnLetter[RandInt(lowerEnLetterLen)]
	}
	return string(b)
}

// RandomUpperStr 随机字符串大写
func RandomUpperStr(l uint32) string {
	return strings.ToUpper(RandomLowerStr(l))
}
