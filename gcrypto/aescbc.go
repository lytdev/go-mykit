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
// @param ivAes 和secretKey位数一致的偏移向量,默认的偏移量就是密钥
func AesCbcEncrypt(plainText, secretKey, ivAes []byte) (cipherText []byte, err error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	paddingText := PKCS5Padding(plainText, block.BlockSize())
	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != block.BlockSize() {
			return nil, ErrIvAes
		} else {
			iv = ivAes
		}
	} else {
		iv = secretKey
	} // To initialize the vector, it needs to be the same length as block.blocksize
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText = make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}

// AesCbcDecrypt
// @param cipherText 加密后的数据
// @param secretKey 密钥
// @param ivAes 和secretKey位数一致的偏移向量,默认的偏移量就是密钥
func AesCbcDecrypt(cipherText, secretKey, ivAes []byte) ([]byte, error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != block.BlockSize() {
			return nil, ErrIvAes
		} else {
			iv = ivAes
		}
	} else {
		iv = secretKey
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	paddingText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(paddingText, cipherText)

	plainText := PKCS5UnPadding(paddingText)
	return plainText, nil
}

func AesCbcEncryptBase64(plainText, secretKey, ivAes []byte) (string, error) {
	encrypted, err := AesCbcEncrypt(plainText, secretKey, ivAes)
	return base64.StdEncoding.EncodeToString(encrypted), err
}

func AesCbcDecryptByBase64(cipherTextBase64 string, secretKey, ivAes []byte) ([]byte, error) {
	plainText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return []byte{}, err
	}
	return AesCbcDecrypt(plainText, secretKey, ivAes)
}

func AesCbcEncryptHex(plainText, secretKey, ivAes []byte) (string, error) {
	encrypted, err := AesCbcEncrypt(plainText, secretKey, ivAes)
	return hex.EncodeToString(encrypted), err
}

func AesCbcDecryptByHex(cipherTextHex string, secretKey, ivAes []byte) ([]byte, error) {
	plainText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return []byte{}, err
	}
	return AesCbcDecrypt(plainText, secretKey, ivAes)
}
