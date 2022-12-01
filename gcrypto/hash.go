package gcrypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

//摘要算法是一种能产生特殊输出格式的算法，这种算法的特点是：无论用户输入什么长度的原始数据，
//经过计算后输出的密文都是固定长度的，这种算法的原理是根据一定的运算规则对原数据进行某种形式的提取，
//这种提取就是摘要，被摘要的数据内容与原数据有密切联系，只要原数据稍有改变，输出的“摘要”便完全不同，
//因此，基于这种原理的算法便能对数据完整性提供较为健全的保障。
//但是，由于输出的密文是提取原数据经过处理的定长值，所以它已经不能还原为原数据，
//即消息摘要算法是不可逆的，理论上无法通过反向运算取得原数据内容，因此它通常只能被用来做数据完整性验证。

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
	_, err := io.WriteString(h, body)
	if err != nil {
		return nil
	}
	return h.Sum(nil)
}

func HmacSha256Hex(key []byte, body string) string {
	return hex.EncodeToString(HmacSha256(key, body))
}

func HmacSha512(key []byte, body string) []byte {
	h := hmac.New(sha512.New, key)
	_, err := io.WriteString(h, body)
	if err != nil {
		return nil
	}
	return h.Sum(nil)
}

func HmacSha512Hex(key []byte, body string) string {
	return hex.EncodeToString(HmacSha512(key, body))
}

func Murmur32Hash(data []byte) uint32 {
	return Sum32(data)
}

func Murmur32HashWithSeed(seed uint32, data []byte) uint32 {
	return Sum32WithSeed(data, seed)
}

func Murmur64Hash(data []byte) uint64 {
	return Sum64(data)
}

func Murmur64HashWithSeed(seed uint32, data []byte) uint64 {
	return Sum64WithSeed(data, seed)
}

func Murmur128Hash(data []byte) (uint64, uint64) {
	return Sum128(data)
}

func Murmur128HashWithSeed(seed uint32, data []byte) (uint64, uint64) {
	return Sum128WithSeed(data, seed)
}
