package gcrypto

import (
	"testing"
)

var (
	eccMsg          = "床前明月光，疑是地上霜，举头望明月，低头思故乡"
	eccBase64PubKey = "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAElJ+LbZBekYTu/Md4T/j3DJsmJFf/3wLLmfUR7sLXCzS1PsDpHIC0QXRdVVdzS9BmP5GdtpesR4Oeh7g0TBBoLA=="
	eccBase64PriKey = "MHcCAQEEIKPH4RlH9IQYwalxykgwlZkV9JjxQW2mHM+oGp4dxkMGoAoGCCqGSM49AwEHoUQDQgAElJ+LbZBekYTu/Md4T/j3DJsmJFf/3wLLmfUR7sLXCzS1PsDpHIC0QXRdVVdzS9BmP5GdtpesR4Oeh7g0TBBoLA=="

	eccHexPubKey = "3059301306072a8648ce3d020106082a8648ce3d030107034200043d39b48322518e8c6053ff63ef0426537fb1d5e16d128802c4c54104d61f84605b6bfa3266cc7f38968c0174d672e3690e50a93c819589f6d0f6bb44a57bcee8"
	eccHexPriKey = "30770201010420af9497e1c61ffe6019592a25f22a12e079e87d935b01bd2dc6d817744053a849a00a06082a8648ce3d030107a144034200043d39b48322518e8c6053ff63ef0426537fb1d5e16d128802c4c54104d61f84605b6bfa3266cc7f38968c0174d672e3690e50a93c819589f6d0f6bb44a57bcee8"
)

func TestEccEncryptBase64(t *testing.T) {
	//公钥加密,私钥解密
	base64Key, err := GenerateEccKeyBase64()
	if err != nil {
		t.Error(err)
	}
	t.Log("**************EccEncryptBase64*******************")
	cipherText1, err := EccEncryptToBase64([]byte(eccMsg), eccBase64PubKey)
	t.Log("公钥加密后(base64)：" + cipherText1)
	plainText1, err := EccDecryptByBase64(cipherText1, eccBase64PriKey)
	t.Log("base64加密串私钥解密后：" + string(plainText1))
	t.Log("*************自动生成公钥和私钥加解密**************")
	cipherText1, err = EccEncryptToBase64([]byte(eccMsg), base64Key.PublicKey)
	t.Log("公钥加密后(base64)：" + cipherText1)
	plainText1, err = EccDecryptByBase64(cipherText1, base64Key.PrivateKey)
	t.Log("base64加密串私钥解密后：" + string(plainText1))

	t.Log("**************EccEncryptHex*******************")
	hexKey, err := GenerateEccKeyHex()
	if err != nil {
		t.Error(err)
	}
	cipherText2, err := EccEncryptToHex([]byte(eccMsg), eccHexPubKey)
	t.Log("公钥加密后(Hex)：" + cipherText1)
	plainText2, err := EccDecryptByHex(cipherText2, eccHexPriKey)
	t.Log("Hex加密串私钥解密后：" + string(plainText2))
	t.Log("************自动生成公钥和私钥加解密************")
	cipherText2, err = EccEncryptToHex([]byte(eccMsg), hexKey.PublicKey)
	t.Log("公钥加密后(Hex)：" + cipherText2)
	plainText1, err = EccDecryptByHex(cipherText2, hexKey.PrivateKey)
	t.Log("Hex加密串私钥解密后：" + string(plainText2))
}

func TestEccSignBase64(t *testing.T) {
	//私钥签名,公钥验签
	base64Key, err := GenerateEccKeyBase64()
	if err != nil {
		t.Error(err)
	}
	rText1, sText1, err := EccSignBase64([]byte(eccMsg), base64Key.PrivateKey)
	rText2, sText2, err := EccSignBase64([]byte(eccMsg), eccBase64PriKey)
	t.Logf("自动私钥签名base64: \r\n rSign:%s \r\n sSign:%s", rText1, sText1)
	t.Logf("手动私钥签名base64: \r\n rSign:%s \r\n sSign:%s", rText2, sText2)
	res1 := EccVerifySignBase64([]byte(eccMsg), rText1, sText1, base64Key.PublicKey)
	t.Logf("自动私钥base64验签结果：%v ", res1)

	res2 := EccVerifySignBase64([]byte(eccMsg), rText2, sText2, eccBase64PubKey)
	t.Logf("自动私钥base64验签结果：%v ", res2)

}

func TestEccSignHex(t *testing.T) {
	//私钥签名,公钥验签
	hex64Key, err := GenerateEccKeyHex()
	if err != nil {
		t.Error(err)
	}
	rText1, sText1, err := EccSignHex([]byte(eccMsg), hex64Key.PrivateKey)
	rText2, sText2, err := EccSignHex([]byte(eccMsg), eccHexPriKey)
	t.Logf("自动私钥签名hex: \r\n rSign:%s \r\n sSign:%s", rText1, sText1)
	t.Logf("手动私钥签名hex: \r\n rSign:%s \r\n sSign:%s", rText2, sText2)
	res1 := EccVerifySignHex([]byte(eccMsg), rText1, sText1, hex64Key.PublicKey)
	t.Logf("自动私钥hex验签结果：%v ", res1)

	res2 := EccVerifySignHex([]byte(eccMsg), rText2, sText2, eccHexPubKey)
	t.Logf("自动私钥hex验签结果：%v ", res2)
}
