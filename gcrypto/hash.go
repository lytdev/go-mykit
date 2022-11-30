package gcrypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

func Sha512Hex(data []byte) string {
	return hex.EncodeToString(Sha512(data))
}

func Sha512(data []byte) []byte {
	digest := sha512.New()
	digest.Write(data)
	return digest.Sum(nil)
}

func Sha1Hex(data []byte) string {
	return hex.EncodeToString(Sha1(data))
}

func Sha1(data []byte) []byte {
	digest := sha1.New()
	digest.Write(data)
	return digest.Sum(nil)
}

func Sha256Hex(data []byte) string {
	return hex.EncodeToString(Sha256(data))
}

func Sha256(data []byte) []byte {
	digest := sha256.New()
	digest.Write(data)
	return digest.Sum(nil)
}

func HmacSha256(key []byte, body string) []byte {
	h := hmac.New(sha256.New, key)
	io.WriteString(h, body)
	return h.Sum(nil)
}

func HmacSha256Hex(key []byte, body string) string {
	return hex.EncodeToString(HmacSha256(key, body))
}

func HmacSha512(key []byte, body string) []byte {
	h := hmac.New(sha512.New, key)
	io.WriteString(h, body)
	return h.Sum(nil)
}

func HmacSha512Hex(key []byte, body string) string {
	return hex.EncodeToString(HmacSha512(key, body))
}
