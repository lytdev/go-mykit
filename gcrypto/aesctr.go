package gcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

// AesCtrEncrypt AES CTR模式的加密
func AesCtrEncrypt(plainText, secretKey []byte) (cipherText []byte, err error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, secretKey[:block.BlockSize()])

	cipherText = make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText, nil
}

// AesCtrDecrypt AES CTR模式的解密
func AesCtrDecrypt(cipherText, secretKey []byte) (plainText []byte, err error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, secretKey[:block.BlockSize()])

	plainText = make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText, nil
}

func AesCtrEncryptBase64(plainText, secretKey []byte) (string, error) {
	encrypted, err := AesCtrEncrypt(plainText, secretKey)
	return base64.StdEncoding.EncodeToString(encrypted), err
}

func AesCtrEncryptHex(plainText, secretKey []byte) (string, error) {
	encrypted, err := AesCtrEncrypt(plainText, secretKey)
	return hex.EncodeToString(encrypted), err
}

func AesCtrDecryptByBase64(cipherTextBase64 string, secretKey []byte) ([]byte, error) {
	plainText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return []byte{}, err
	}
	return AesCtrDecrypt(plainText, secretKey)
}

func AesCtrDecryptByHex(cipherTextHex string, secretKey []byte) ([]byte, error) {
	plainText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return []byte{}, err
	}
	return AesCtrDecrypt(plainText, secretKey)
}
