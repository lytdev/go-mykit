package gcrypto

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestEncryptByAes(t *testing.T) {
	origData := []byte("这是一个待加密的明文密码") // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP")  // 加密的密钥
	t.Log("原文：", string(origData))

	t.Log("------------------ 默认加密 --------------------")
	encrypted1, _ := EncryptDataByAes("ABCDEFGHIJKLMNOP", origData)
	t.Log("密文(base64)：", encrypted1)
	decrypted1, _ := DecryptDataByAes("ABCDEFGHIJKLMNOP", encrypted1)
	t.Log("解密结果：", string(decrypted1))

	t.Log("------------------ CBC模式 --------------------")
	encrypted, _ := AesEncryptCBC(key, origData)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted, _ := AesDecryptCBC(key, encrypted)
	t.Log("解密结果：", string(decrypted))

	t.Log("------------------ ECB模式 --------------------")
	encrypted = AesEncryptECB(key, origData)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptECB(key, encrypted)
	t.Log("解密结果：", string(decrypted))

	t.Log("------------------ CFB模式 --------------------")
	encrypted = AesEncryptCFB(key, origData)
	t.Log("密文(hex)：", hex.EncodeToString(encrypted))
	t.Log("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptCFB(key, encrypted)
	t.Log("解密结果：", string(decrypted))
}
