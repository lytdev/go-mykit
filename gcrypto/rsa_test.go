package gcrypto

import (
	"encoding/base64"
	"testing"
)

func TestRsa(t *testing.T) {
	bits := 2048
	spub, spri := CreateKeyX509PKCS1(bits)
	data := "这是一个待加密的字符串文本数据"
	pub, _ := PublicKeyFromX509PKCS1(spub)

	pri, _ := PrivateKeyFromX509PKCS1(spri)

	encrypt, _ := RsaEncrypt(pub, []byte(data))
	pwd1 := base64.StdEncoding.EncodeToString(encrypt)
	t.Log("密文(base64)：", pwd1)
	tmp, _ := base64.StdEncoding.DecodeString(pwd1)
	decrypt, _ := RsaDecrypt(pri, tmp)

	t.Log("解密结果：", string(decrypt))
}

func TestSign(t *testing.T) {
	bits := 2048
	spub, spri := CreateKeyX509PKCS1(bits)
	data := "这是一段待签名的字符串数据"
	pub, err := PublicKeyFromX509PKCS1(spub)

	t.Log(spub, spri)
	pri, err := PrivateKeyFromX509PKCS1(spri)

	sign, err := Sign(pri, []byte(data))
	if err != nil {
		return
	}
	err = Verify(pub, sign, []byte(data))

}
