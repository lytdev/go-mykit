package gcrypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
)

// ECC椭圆曲线
// GO里面只有ECC数字签名的接口，所以我们这里实现了ECC的数字签名功能，ECC椭圆曲线加密使用了区块链以太坊中的相关接口,ECC一般只签名使用加密一般不使用

// eccEncrypt ECC使用公钥进行加密
func eccEncrypt(plainText, pubKey []byte) (cipherText []byte, err error) {
	tempPublicKey, err := x509.ParsePKIXPublicKey(pubKey)
	if err != nil {
		return nil, err
	}
	// Decode to get the private key in the ecdsa package
	publicKey1 := tempPublicKey.(*ecdsa.PublicKey)
	// Convert to the public key in the ecies package in the ethereum package
	publicKey := ImportECDSAPublic(publicKey1)
	cipherText, err = Encrypt(rand.Reader, publicKey, plainText, nil, nil)
	return cipherText, err

}

// eccDecrypt ECC使用私钥进行加密
func eccDecrypt(cipherText, priKey []byte) (msg []byte, err error) {
	tempPrivateKey, err := x509.ParseECPrivateKey(priKey)
	if err != nil {
		return nil, err
	}
	// Decode to get the private key in the ecdsa package
	// Convert to the private key in the ecies package in the ethereum package
	privateKey := ImportECDSA(tempPrivateKey)
	plainText, err := privateKey.Decrypt(cipherText, nil, nil)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// EccEncryptToBase64 ECC使用公钥进行加密后转base64
func EccEncryptToBase64(plainText []byte, base64PubKey string) (base64CipherText string, err error) {
	pub, err := base64.StdEncoding.DecodeString(base64PubKey)
	if err != nil {
		return "", err
	}
	cipherBytes, err := eccEncrypt(plainText, pub)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// EccDecryptByBase64 ECC使用私钥对加密后的base64字符串进行解密
func EccDecryptByBase64(base64CipherText, base64PriKey string) (plainText []byte, err error) {
	privateBytes, err := base64.StdEncoding.DecodeString(base64PriKey)
	if err != nil {
		return nil, err
	}
	cipherTextBytes, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return nil, err
	}
	return eccDecrypt(cipherTextBytes, privateBytes)
}

// EccEncryptToHex ECC使用公钥进行加密后转Hex
func EccEncryptToHex(plainText []byte, hexPubKey string) (hexCipherText string, err error) {
	pub, err := hex.DecodeString(hexPubKey)
	if err != nil {
		return "", err
	}
	cipherBytes, err := eccEncrypt(plainText, pub)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherBytes), nil
}

// EccDecryptByHex ECC使用私钥对加密后的Hex字符串进行解密
func EccDecryptByHex(hexCipherText, hexPriKey string) (plainText []byte, err error) {
	privateBytes, err := hex.DecodeString(hexPriKey)
	if err != nil {
		return nil, err
	}
	cipherTextBytes, err := hex.DecodeString(hexCipherText)
	if err != nil {
		return nil, err
	}
	return eccDecrypt(cipherTextBytes, privateBytes)
}
