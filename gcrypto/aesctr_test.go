package gcrypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ctrPlaintext = "飞流直下三千尺，疑似银河落九天。"
)

func TestAesCtrEncryptAndDecrypt(t *testing.T) {
	origData := []byte(ctrPlaintext) // 待加密的数据
	key := []byte(key16)             // 加密的密钥
	t.Log("密钥：", string(key))
	t.Log("原文：", string(origData))

	t.Log("******************进行ECB加密*****************")
	encrypted1, err1 := AesCtrEncryptBase64(origData, key, nil)
	if err1 != nil {
		t.Error(err1)
	}
	encrypted2, err1 := AesCtrEncryptHex(origData, key, nil)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("密文(base64)：", encrypted1)
	t.Log("密文(Hex)：", encrypted2)
	t.Log("****************进行ECB解密*****************")
	decrypted1, err1 := AesCtrDecryptByBase64(encrypted1, key, nil)
	if err1 != nil {
		t.Error(err1)
	}
	decrypted2, err1 := AesCtrDecryptByHex(encrypted2, key, nil)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("base64解密结果：", string(decrypted1))
	t.Log("Hex解密结果：", string(decrypted2))
	// 密钥： FgTyct3gH9QfWnTh
	// 原文： 飞流直下三千尺，疑似银河落九天。
	// ******************进行ECB加密*****************
	// 密文(base64)： wzsUNdeUUlj88to4cBC4KW1PBdNw/H0bp7gUluqMPcNaLoU51uJUnG55rCjLon6D
	// 密文(Hex)： c33b1435d7945258fcf2da387010b8296d4f05d370fc7d1ba7b81496ea8c3dc35a2e8539d6e2549c6e79ac28cba27e83
	// ****************进行ECB解密*****************
	// base64解密结果： 飞流直下三千尺，疑似银河落九天。
	// Hex解密结果： 飞流直下三千尺，疑似银河落九天。
}

func TestAesCtr(t *testing.T) {
	cipherBytes, err := AesCtrEncrypt([]byte(ctrPlaintext), []byte(key16), nil)
	assert.Nil(t, err)
	text, err := AesCtrDecrypt(cipherBytes, []byte(key16), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), ctrPlaintext)

	cipherBytes, err = AesCtrEncrypt([]byte(ctrPlaintext), []byte(key24), nil)
	assert.Nil(t, err)
	text, err = AesCtrDecrypt(cipherBytes, []byte(key24), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), ctrPlaintext)

	cipherBytes, err = AesCtrEncrypt([]byte(ctrPlaintext), []byte(key32), nil)
	assert.Nil(t, err)
	text, err = AesCtrDecrypt(cipherBytes, []byte(key32), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), ctrPlaintext)

	cipherBytes, err = AesCtrEncrypt([]byte(ctrPlaintext), []byte(key), nil)
	assert.NotNil(t, err)
	text, err = AesCtrDecrypt(cipherBytes, []byte(key), nil)
	assert.NotNil(t, err)
}

func TestAesCtrEncryptBase64(t *testing.T) {
	cipher, err := AesCtrEncryptBase64([]byte(ctrPlaintext), []byte(key16), nil)
	assert.Nil(t, err)
	text, err := AesCtrDecryptByBase64(cipher, []byte(key16), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), ctrPlaintext)

	_, err = AesCtrDecryptByBase64("11111", []byte(key16), nil)
	assert.NotNil(t, err)
}

func TestAesCtrEncryptHex(t *testing.T) {
	cipher, err := AesCtrEncryptHex([]byte(ctrPlaintext), []byte(key16), nil)
	assert.Nil(t, err)
	text, err := AesCtrDecryptByHex(cipher, []byte(key16), nil)
	assert.Nil(t, err)
	assert.Equal(t, string(text), ctrPlaintext)

	_, err = AesCtrDecryptByHex("11111", []byte(key16), nil)
	assert.NotNil(t, err)
}
