package gcrypto

import (
	"bytes"
	"errors"
	"fmt"
)

// pkcs7Padding 填充
func pkcs7Padding(src []byte, blockSize int) (dest []byte, err error) {
	if blockSize <= 0 {
		return nil, errors.New("block size is 0")
	} else if src == nil || len(src) == 0 {
		return nil, errors.New("src is nil")
	}
	//判断缺少几位长度,最少1,最多blockSize
	n := blockSize - (len(src) % blockSize)
	pb := make([]byte, len(src)+n)
	copy(pb, src)
	//补足位数,把切片[]byte{byte(padding)}复制padding个
	copy(pb[len(src):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(src []byte, blockSize int) (dest []byte, err error) {

	if blockSize <= 0 {
		return nil, errors.New("block size is 0")
	} else if len(src)%blockSize != 0 {
		return nil, errors.New("src length error")
	} else if src == nil || len(src) == 0 {
		return nil, errors.New("src is nil")
	}
	c := src[len(src)-1]
	//获取填充的个数
	padLength := int(c)
	if padLength == 0 || padLength > len(src) {
		return nil, errors.New("pad length error")
	}
	for i := 0; i < padLength; i++ {
		if src[len(src)-padLength+i] != c {
			return nil, errors.New("pad content error")
		}
	}
	return src[:len(src)-padLength], nil

}

// pkcs5Padding PKCS5填充
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pkcs5UnPadding 去除PKCS5填充
func pkcs5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	up := int(origData[length-1])

	if length < up {
		return nil, fmt.Errorf("invalid unpadding length")
	}
	return origData[:(length - up)], nil
}
