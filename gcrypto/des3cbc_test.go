package gcrypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDes3EncryptAndDecrypt(t *testing.T) {
	cipher, _ := TripleDesEncryptBase64([]byte(plaintext), []byte(key24), nil)
	t.Logf("3重DES加密后数据:" + cipher)
	text, _ := TripleDesDecryptByBase64(cipher, []byte(key24), nil)
	t.Logf("3重DES解密后数据:" + string(text))
}

func TestTripleDesCbc(t *testing.T) {
	cipherBytes, err := TripleDesEncrypt([]byte(plaintext), []byte(key24), nil)
	assert.Nil(t, err)
	text, err := TripleDesDecrypt(cipherBytes, []byte(key24), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)
	assert.NotEqual(t, string(text), "test")

	cipherBytes, err = TripleDesEncrypt([]byte(plaintext), []byte(key24), []byte(desGoodiv))
	assert.Nil(t, err)
	text, err = TripleDesDecrypt(cipherBytes, []byte(key24), []byte(desGoodiv))
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)
	assert.NotEqual(t, string(text), "test")

	cipherBytes, err = TripleDesEncrypt([]byte(plaintext), []byte(key24), []byte(desBadiv))
	assert.NotNil(t, err)

	cipherBytes, err = TripleDesEncrypt([]byte(plaintext), []byte(desBadiv), []byte(desBadiv))
	assert.NotNil(t, err)

}

func TestTripleDesEncryptBase64(t *testing.T) {
	cipher, err := TripleDesEncryptBase64([]byte(plaintext), []byte(key24), nil)
	assert.Nil(t, err)
	text, err := TripleDesDecryptByBase64(cipher, []byte(key24), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	_, err = TripleDesDecryptByBase64("11111", []byte(key24), nil)
	assert.NotNil(t, err)
}

func TestTripleDesEncryptHex(t *testing.T) {
	cipher, err := TripleDesEncryptHex([]byte(plaintext), []byte(key24), nil)
	assert.Nil(t, err)
	text, err := TripleDesDecryptByHex(cipher, []byte(key24), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	_, err = TripleDesDecryptByHex("11111", []byte(key24), nil)
	assert.NotNil(t, err)
}
