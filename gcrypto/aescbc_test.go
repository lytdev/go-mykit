package gcrypto

import (
	"testing"
)

func TestAesCbc(t *testing.T) {
	origData := []byte("飞流直下三千尺，疑似银河落九天。") // 待加密的数据
	key := []byte("FgTyct3gH9QfWnTh")      // 加密的密钥
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

func TestAesCbcEncrypt(t *testing.T) {
	origData := []byte("飞流直下三千尺，疑似银河落九天。") // 待加密的数据
	key := []byte("FgTyct3gH9QfWnTh")      // 加密的密钥
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
}

func TestAesCbcDecrypt(t *testing.T) {
	decrypted1 := "YHTiv3efNoH+FPuDI58ZRdATIdDR7hBdje6Gl5Guk3+v//i/IVwVJ9Zd+T5sPURsVA6f4nt7CrkQdGWp6TFvYA=="                                         // 待解密的数据
	decrypted2 := "6074e2bf779f3681fe14fb83239f1945d01321d0d1ee105d8dee869791ae937faffff8bf215c1527d65df93e6c3d446c540e9fe27b7b0ab9107465a9e9316f60" // 待解密的数据
	key := []byte("FgTyct3gH9QfWnTh")                                                                                                                // 加密的密钥
	t.Log("****************进行CBC解密******************")
	plaintext1, err1 := AesCbcDecryptByBase64(decrypted1, key, key)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("base64解密结果：", string(plaintext1))
	plaintext2, err1 := AesCbcDecryptByHex(decrypted2, key, key)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("Hex解密结果：", string(plaintext2))
}
