package gcrypto

import (
	"testing"
)

func TestAesCtrEncrypt(t *testing.T) {
	origData := []byte("这是一个待加密的明文密码") // 待加密的数据
	key := []byte("FgTyct3gH9QfWnTh")  // 加密的密钥
	t.Log("原文：", string(origData))

	t.Log("******************CTR加密*****************")
	encrypted1, err1 := AesCtrEncryptBase64(origData, key)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("密文(base64)：", encrypted1)
	t.Log("******************************************************")
	decrypted1, err2 := AesCtrDecryptByBase64(encrypted1, key)
	if err2 != nil {
		t.Error(err2)
	}
	t.Log("解密结果：", string(decrypted1))

}
