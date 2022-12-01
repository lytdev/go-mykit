package gcrypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	plaintext = "飞流直下三千尺，疑似银河落九天。"
	key       = "111"
	key16     = "FgTyct3gH9QfWnTh"
	aesBadiv  = "FgTyct3gH9QfWn"
	aesGoodiv = "FgTyct3gH9QfWnTh"
	key24     = "c7ANgV8z67VcPmxXfYjZxqEN"
	key32     = "VddYB8XGvvGSGfuUcDHUvX3nnMEKCh49"
)

func TestAesCbcEncryptAndDecrypt(t *testing.T) {
	origData := []byte(plaintext) // 待加密的数据
	key := []byte(key16)          // 加密的密钥
	t.Log("密钥：", string(key))
	t.Log("原文：", string(origData))

	t.Log("****************进行CBC加密*****************")
	encrypted1, err1 := AesCbcEncryptBase64(origData, key, key)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("加密后密文(base64)：", encrypted1)
	encrypted2, err1 := AesCbcEncryptHex(origData, key, key)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("加密后密文(Hex)：", encrypted2)
	t.Log("****************进行CBC解密******************")
	decrypted1, err1 := AesCbcDecryptByBase64(encrypted1, key, key)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("base64解密结果：", string(decrypted1))
	decrypted2, err1 := AesCbcDecryptByHex(encrypted2, key, key)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("Hex解密结果：", string(decrypted2))
	// 密钥： FgTyct3gH9QfWnTh
	// 原文： 飞流直下三千尺，疑似银河落九天。
	// ****************进行CBC加密*****************
	// 加密后密文(base64)： YHTiv3efNoH+FPuDI58ZRdATIdDR7hBdje6Gl5Guk3+v//i/IVwVJ9Zd+T5sPURsVA6f4nt7CrkQdGWp6TFvYA==
	// 加密后密文(Hex)： 6074e2bf779f3681fe14fb83239f1945d01321d0d1ee105d8dee869791ae937faffff8bf215c1527d65df93e6c3d446c540e9fe27b7b0ab9107465a9e9316f60
	// ****************进行CBC解密******************
	// base64解密结果： 飞流直下三千尺，疑似银河落九天。
	// Hex解密结果： 飞流直下三千尺，疑似银河落九天。
}

func TestAesCbc(t *testing.T) {

	cipherBytes, err := AesCbcEncrypt([]byte(plaintext), []byte(key16), nil)
	assert.Nil(t, err)
	text, err := AesCbcDecrypt(cipherBytes, []byte(key16), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	_, err = AesCbcDecrypt(cipherBytes, []byte(key24), nil)
	assert.NotNil(t, err)
	text, err = AesCbcDecrypt([]byte("badtext"), []byte(key24), nil)
	assert.Equal(t, string(text), "")

	cipherBytes, err = AesCbcEncrypt([]byte(plaintext), []byte(key24), nil)
	assert.Nil(t, err)
	text, err = AesCbcDecrypt(cipherBytes, []byte(key24), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	cipherBytes, err = AesCbcEncrypt([]byte(plaintext), []byte(key32), nil)
	assert.Nil(t, err)
	text, err = AesCbcDecrypt(cipherBytes, []byte(key32), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	cipherBytes, err = AesCbcEncrypt([]byte(plaintext), []byte(key), nil)
	assert.NotNil(t, err)
	text, err = AesCbcDecrypt(cipherBytes, []byte(key), nil)
	assert.NotNil(t, err)

	cipherBytes, err = AesCbcEncrypt([]byte(plaintext), []byte(key16), []byte(aesBadiv))
	assert.NotNil(t, err)
	text, err = AesCbcDecrypt(cipherBytes, []byte(key16), []byte(aesBadiv))
	assert.NotNil(t, err)

	cipherBytes, err = AesCbcEncrypt([]byte(plaintext), []byte(key16), []byte(aesGoodiv))
	assert.Nil(t, err)
	text, err = AesCbcDecrypt(cipherBytes, []byte(key16), []byte(aesGoodiv))
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)
}

func TestAesCbcEncryptBase64(t *testing.T) {
	cipher, err := AesCbcEncryptBase64([]byte(plaintext), []byte(key16), nil)
	assert.Nil(t, err)
	text, err := AesCbcDecryptByBase64(cipher, []byte(key16), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	_, err = AesCbcDecryptByBase64("11111", []byte(key16), nil)
	assert.NotNil(t, err)
}

func TestAesCbcEncryptHex(t *testing.T) {
	cipher, err := AesCbcEncryptHex([]byte(plaintext), []byte(key16), nil)
	assert.Nil(t, err)
	text, err := AesCbcDecryptByHex(cipher, []byte(key16), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), plaintext)

	_, err = AesCbcDecryptByHex("11111", []byte(key16), nil)
	assert.NotNil(t, err)
}
