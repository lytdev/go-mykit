package gcrypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	key8      = "12345678"
	desBadiv  = "11111"
	desGoodiv = "12345678"
)

func TestDesCbcDecryptAndDecrypt(t *testing.T) {
	cipher, err := DesCbcEncryptBase64([]byte(plaintext), []byte(key8), nil)
	assert.Nil(t, err)
	t.Log("加密后数据:" + cipher)
	text, err := DesCbcDecryptByBase64(cipher, []byte(key8), nil)
	t.Log("解密后数据:" + string(text))
}

func TestDesCbc(t *testing.T) {

	cipherBytes, err := DesCbcEncrypt([]byte(plaintext), []byte(key8), nil)
	assert.Nil(t, err)
	text, err := DesCbcDecrypt(cipherBytes, []byte(key8), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	cipherBytes, err = DesCbcEncrypt([]byte(plaintext), []byte(key), nil)
	assert.NotNil(t, err)

	cipherBytes, err = DesCbcEncrypt([]byte(plaintext), []byte(key8), []byte(desBadiv))
	assert.NotNil(t, err)

	cipherBytes, err = DesCbcEncrypt([]byte(plaintext), []byte(key8), []byte(desGoodiv))
	assert.Nil(t, err)
	text, err = DesCbcDecrypt(cipherBytes, []byte(key8), []byte(desGoodiv))
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

}

func TestDesEncryptBase64(t *testing.T) {
	cipher, err := DesCbcEncryptBase64([]byte(plaintext), []byte(key8), nil)
	assert.Nil(t, err)
	text, err := DesCbcDecryptByBase64(cipher, []byte(key8), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	_, err = DesCbcDecryptByBase64("11111", []byte(key8), nil)
	assert.NotNil(t, err)
}

func TestDesEncryptHex(t *testing.T) {
	cipher, err := DesCbcEncryptHex([]byte(plaintext), []byte(key8), nil)
	assert.Nil(t, err)
	text, err := DesCbcDecryptByHex(cipher, []byte(key8), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	_, err = DesCbcDecryptByHex("11111", []byte(key8), nil)
	assert.NotNil(t, err)
}
