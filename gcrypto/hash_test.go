package gcrypto

import (
	"github.com/lytdev/go-mykit/gconv"
	"testing"
)

func TestSha1Hex(t *testing.T) {
	actual := Sha1Hex([]byte(plaintext))
	t.Log("Sha1Hex：" + actual)
}

func TestSha256Hex(t *testing.T) {
	actual := Sha256Hex([]byte(plaintext))
	t.Log("Sha256Hex：" + actual)
}

func TestSha512Hex(t *testing.T) {
	actual := Sha512Hex([]byte(plaintext))
	t.Log("Sha512Hex：" + actual)
}

func TestMurmur3(t *testing.T) {
	v1 := Murmur32Hash([]byte(plaintext))
	t.Log("Murmur32Hash：" + gconv.ToString(v1))
	v2 := Murmur32HashWithSeed(1, []byte(plaintext))
	t.Log("Murmur32HashWithSeed：" + gconv.ToString(v2))
	v3 := Murmur64Hash([]byte(plaintext))
	t.Log("Murmur64Hash：" + gconv.ToString(v3))
	v4 := Murmur64HashWithSeed(1, []byte(plaintext))
	t.Log("Murmur64HashWithSeed：" + gconv.ToString(v4))
	v51, v52 := Murmur128Hash([]byte(plaintext))
	t.Log("Murmur128Hash：" + gconv.ToString(v51))
	t.Log("Murmur128Hash：" + gconv.ToString(v52))
	v61, v62 := Murmur128HashWithSeed(1, []byte(plaintext))
	t.Log("Murmur128HashWithSeed：" + gconv.ToString(v61))
	t.Log("Murmur128HashWithSeed：" + gconv.ToString(v62))
}

func BenchmarkSha256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sha256Hex([]byte(plaintext))
	}
}

func BenchmarkSha512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sha512Hex([]byte(plaintext))
	}
}

func TestHmacSha256Hex(t *testing.T) {
	res := HmacSha256Hex([]byte("test"), plaintext)
	t.Log("HmacSha256Hex：" + res)
}

func TestHmacSha512Hex(t *testing.T) {
	res := HmacSha512Hex([]byte("test"), plaintext)
	t.Log("HmacSha512Hex：" + res)
}
