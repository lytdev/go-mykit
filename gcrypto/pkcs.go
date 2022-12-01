package gcrypto

import (
	"bytes"
)

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padText...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func PKCS5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func PKCS5UnPadding(plainText []byte, blockSize int) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length || number > blockSize {
		return nil, ErrPaddingSize
	}
	return plainText[:length-number], nil
}
