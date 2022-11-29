package grandom

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// RandomString 生成随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(t))
		} else if t < 36 {
			result = append(result, string(rune(rand.Intn(26)+65)))
		} else {
			result = append(result, string(rune(rand.Intn(26)+97)))
		}
	}
	return strings.Join(result, "")
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
