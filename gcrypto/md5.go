package gcrypto

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MD5Upper MD5大写加密32位
func MD5Upper(value string) string {
	return strings.ToUpper(Md5(value))
}

// MD5To16Upper MD5大写加密16位
func MD5To16Upper(value string) string {
	return strings.ToUpper(Md5To16(value)[8:24])
}

// Md5 32位md5加密后字符串
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5To16 16位md5加密后字符串
func Md5To16(str string) string {
	return Md5(str)[8:24]
}

// Md5WithSalt 加盐加密
func Md5WithSalt(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// Md5SaltMulti 加盐并多次加密
func Md5SaltMulti(str string, salt string, iteration int) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先传入盐值
	h.Write(b)
	var res []byte
	res = h.Sum(nil)
	for i := 0; i < iteration-1; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	return hex.EncodeToString(res)
}
