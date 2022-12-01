package gcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"github.com/lytdev/go-mykit/gnum"
)

// https://www.cnblogs.com/happyhippy/archive/2006/12/23/601353.html

// AesCbcEncrypt CBC 密文分组链接模式,可以避免ECB模式的缺陷
// @param plainText 待加密的数据
// @param secretKey 密钥
// @param ivAes 和secretKey位数一致的偏移向量,默认的偏移量就是密钥
func AesCbcEncrypt(plainText, secretKey, ivAes []byte) (cipherText []byte, err error) {
	if len(secretKey) != KeyLength16 && len(secretKey) != KeyLength24 && len(secretKey) != KeyLength32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	bLen := block.BlockSize()
	paddingText := PKCS5Padding(plainText, bLen)
	ptLen := len(paddingText)
	if !gnum.NumMulti(ptLen, bLen) {
		return nil, ErrAesSrcBlockSize
	}
	var iv []byte
	// To initialize the vector, it needs to be the same length as block.BlockSize
	if len(ivAes) != 0 {
		if len(ivAes) != bLen {
			return nil, ErrIvAes
		} else {
			iv = ivAes
		}
	} else {
		iv = secretKey[:bLen]
	}

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
	if len(secretKey) != KeyLength16 && len(secretKey) != KeyLength24 && len(secretKey) != KeyLength32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	bLen := block.BlockSize()
	ctLen := len(cipherText)
	if !gnum.NumMulti(ctLen, bLen) {
		return nil, ErrAesSrcBlockSize
	}
	var iv []byte
	// To initialize the vector, it needs to be the same length as block.BlockSize
	if len(ivAes) != 0 {
		if len(ivAes) != bLen {
			return nil, ErrIvAes
		} else {
			iv = ivAes
		}
	} else {
		iv = secretKey[:bLen]
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	paddingText := make([]byte, ctLen)
	blockMode.CryptBlocks(paddingText, cipherText)
	return PKCS5UnPadding(paddingText, blockMode.BlockSize())
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
