package gcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

// AesCbcEncrypt
// @param plainText 待加密的数据
// @param secretKey 密钥
func AesCbcEncrypt(plainText, secretKey []byte) (cipherText []byte, err error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	paddingText := PKCS5Padding(plainText, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, secretKey[:block.BlockSize()])
	cipherText = make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}

const Ivaes = "12345678"

// AesCbcDecrypt
// @param cipherText 加密后的数据
// @param secretKey 密钥
func AesCbcDecrypt(cipherText, secretKey []byte) ([]byte, error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, secretKey[:block.BlockSize()])
	paddingText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(paddingText, cipherText)

	plainText := PKCS5UnPadding(paddingText)
	return plainText, nil
}

func AesCbcEncryptBase64(plainText, secretKey []byte) (string, error) {
	encrypted, err := AesCbcEncrypt(plainText, secretKey)
	return base64.StdEncoding.EncodeToString(encrypted), err
}

func AesCbcDecryptByBase64(cipherTextBase64 string, secretKey []byte) ([]byte, error) {
	plainText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return []byte{}, err
	}
	return AesCbcDecrypt(plainText, secretKey)
}

func AesCbcEncryptHex(plainText, secretKey []byte) (string, error) {
	encrypted, err := AesCbcEncrypt(plainText, secretKey)
	return hex.EncodeToString(encrypted), err
}

func AesCbcDecryptByHex(cipherTextHex string, secretKey []byte) ([]byte, error) {
	plainText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return []byte{}, err
	}
	return AesCbcDecrypt(plainText, secretKey)
}
