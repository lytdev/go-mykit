package gcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

// https://www.cnblogs.com/happyhippy/archive/2006/12/23/601353.html

// AesCtrEncrypt AES CTR(计数器模式)模式的加密,NoPadding
// @param plainText 代加密的数据
// @param secretKey 密钥
// @param ivAes 和secretKey位数一致的偏移向量,默认的偏移量就是密钥
func AesCtrEncrypt(plainText, secretKey, ivAes []byte) (cipherText []byte, err error) {
	if len(secretKey) != KeyLength16 && len(secretKey) != KeyLength24 && len(secretKey) != KeyLength32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	bLen := block.BlockSize()
	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != bLen {
			return nil, ErrIvAes
		} else {
			iv = ivAes
		}
	} else {
		iv = secretKey[:bLen]
	}
	stream := cipher.NewCTR(block, iv)

	cipherText = make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText, nil
}

// AesCtrDecrypt AES CTR模式的解密,NoPadding
// @param cipherText 加密后的数据
// @param secretKey 密钥
// @param ivAes 和secretKey位数一致的偏移向量,默认的偏移量就是密钥
func AesCtrDecrypt(cipherText, secretKey, ivAes []byte) (plainText []byte, err error) {
	if len(secretKey) != KeyLength16 && len(secretKey) != KeyLength24 && len(secretKey) != KeyLength32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	bLen := block.BlockSize()
	var iv []byte
	if len(ivAes) != 0 {
		if len(ivAes) != bLen {
			return nil, ErrIvAes
		} else {
			iv = ivAes
		}
	} else {
		iv = secretKey[:bLen]
	}
	stream := cipher.NewCTR(block, iv)

	plainText = make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText, nil
}

func AesCtrEncryptBase64(plainText, secretKey, ivAes []byte) (string, error) {
	encrypted, err := AesCtrEncrypt(plainText, secretKey, ivAes)
	return base64.StdEncoding.EncodeToString(encrypted), err
}

func AesCtrEncryptHex(plainText, secretKey, ivAes []byte) (string, error) {
	encrypted, err := AesCtrEncrypt(plainText, secretKey, ivAes)
	return hex.EncodeToString(encrypted), err
}

func AesCtrDecryptByBase64(cipherTextBase64 string, secretKey, ivAes []byte) ([]byte, error) {
	plainText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return []byte{}, err
	}
	return AesCtrDecrypt(plainText, secretKey, ivAes)
}

func AesCtrDecryptByHex(cipherTextHex string, secretKey, ivAes []byte) ([]byte, error) {
	plainText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return []byte{}, err
	}
	return AesCtrDecrypt(plainText, secretKey, ivAes)
}
