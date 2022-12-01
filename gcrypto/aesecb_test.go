package gcrypto

import "testing"

func TestAesEcb(t *testing.T) {
	origData := []byte("飞流直下三千尺，疑似银河落九天。") // 待加密的数据
	key := []byte("FgTyct3gH9QfWnTh")      // 加密的密钥
	t.Log("密钥：", string(key))
	t.Log("原文：", string(origData))

	t.Log("******************进行ECB加密*****************")
	encrypted1, err1 := AesEcbEncryptBase64(origData, key)
	if err1 != nil {
		t.Error(err1)
	}
	encrypted2, err1 := AesEcbEncryptHex(origData, key)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("密文(base64)：", encrypted1)
	t.Log("密文(Hex)：", encrypted2)
	t.Log("****************进行ECB解密*****************")
	decrypted1, err1 := AesEcbDecryptByBase64(encrypted1, key)
	if err1 != nil {
		t.Error(err1)
	}
	decrypted2, err1 := AesEcbDecryptByHex(encrypted2, key)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log("base64解密结果：", string(decrypted1))
	t.Log("Hex解密结果：", string(decrypted2))
	//密钥： FgTyct3gH9QfWnTh
	//原文： 飞流直下三千尺，疑似银河落九天。
	// ******************进行ECB加密*****************
	// 密文(base64)： VDvc5kiAmutQNR1zgT7Ow0pOQtkM880YwJVcBNvSLSdT01pYgeSPhwNKof6IxjZEwyisuZueWiPtv0g1Cx5n6Q==
	// 密文(Hex)： 543bdce648809aeb50351d73813ecec34a4e42d90cf3cd18c0955c04dbd22d2753d35a5881e48f87034aa1fe88c63644c328acb99b9e5a23edbf48350b1e67e9
	// ****************进行ECB解密*****************
	// base64解密结果： 飞流直下三千尺，疑似银河落九天。
	// Hex解密结果： 飞流直下三千尺，疑似银河落九天。
}
