package gcrypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	shaData = "sha text"

	sha1Hex   = "d61babc269a1ccf83d8d08583fdf513eedeb55e6"
	sha256Hex = "65294d857d822c8b73af70c78cf6fc4325a0bf28c2efbbcd07c55b19eaf20d20"
	sha512Hex = "b4d5cda7b08feeca4ce2bf17e1ffab7d13e5234faca54ae46f4f87f66200a3bbc07b4b37b095eaf3bca2f8dba707bc259af3fe6e6e0b925a43915c9f351d92be"
)

func TestSha1Hex(t *testing.T) {
	actual := Sha1Hex([]byte(shaData))
	assert.Equal(t, sha1Hex, actual)
}

func TestSha256Hex(t *testing.T) {
	actual := Sha256Hex([]byte(shaData))
	assert.Equal(t, sha256Hex, actual)
}

func TestSha512Hex(t *testing.T) {
	actual := Sha512Hex([]byte(shaData))
	assert.Equal(t, sha512Hex, actual)
}

func BenchmarkSha256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sha256Hex([]byte(shaData))
	}
}

func BenchmarkSha512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sha512Hex([]byte(shaData))
	}
}

var (
	hmac256 = "ee59ba4d767a2122cd5acebc538b31adce6d378719872e4b2b961d26a87b01cd"
	hmac512 = "ed697fcce5cc626037d96f8f27fe86f4bfbf12ddf1236b9f4a9172acab1d6c02fb987bb453c525f6dbc0167164e7ac18fcf8d36ed09ece8a7a03222473f57363"
)

func TestHmacSha256Hex(t *testing.T) {
	res := HmacSha256Hex([]byte("test"), "hmac text")
	assert.Equal(t, res, hmac256)
}

func TestHmacSha512Hex(t *testing.T) {
	res := HmacSha512Hex([]byte("test"), "hmac text")
	assert.Equal(t, res, hmac512)
}
