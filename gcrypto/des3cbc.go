package gcrypto

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"github.com/lytdev/go-mykit/gnum"
)

func TripleDesEncrypt(plainText, secretKey, ivDes []byte) ([]byte, error) {
	if len(secretKey) != KeyLength24 {
		return nil, ErrKeyLengthTwentyFour
	}
	block, err := des.NewTripleDESCipher(secretKey)
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
	if len(ivDes) != 0 {
		if len(ivDes) != bLen {
			return nil, ErrIvDes
		} else {
			iv = ivDes
		}
	} else {
		iv = secretKey[:bLen]
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte, ptLen)
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}

func TripleDesDecrypt(cipherText, secretKey, ivDes []byte) ([]byte, error) {
	if len(secretKey) != KeyLength24 {
		return nil, ErrKeyLengthTwentyFour
	}
	// 1. Specifies that the 3des decryption algorithm creates and returns a cipher.Block interface using the TDEA algorithm。
	block, err := des.NewTripleDESCipher(secretKey)
	if err != nil {
		return nil, err
	}
	bLen := block.BlockSize()
	ctLen := len(cipherText)
	if !gnum.NumMulti(ctLen, bLen) {
		return nil, ErrAesSrcBlockSize
	}
	var iv []byte
	if len(ivDes) != 0 {
		if len(ivDes) != bLen {
			return nil, ErrIvDes
		} else {
			iv = ivDes
		}
	} else {
		iv = secretKey[:bLen]
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	paddingText := make([]byte, ctLen)
	blockMode.CryptBlocks(paddingText, cipherText)
	return PKCS5UnPadding(paddingText, bLen)
}

func TripleDesEncryptBase64(plainText, secretKey, ivAes []byte) (string, error) {
	encrypted, err := TripleDesEncrypt(plainText, secretKey, ivAes)
	return base64.StdEncoding.EncodeToString(encrypted), err
}

func TripleDesEncryptHex(plainText, secretKey, ivAes []byte) (string, error) {
	encrypted, err := TripleDesEncrypt(plainText, secretKey, ivAes)
	return hex.EncodeToString(encrypted), err
}

func TripleDesDecryptByBase64(cipherTextBase64 string, secretKey, ivAes []byte) ([]byte, error) {
	plainText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return []byte{}, err
	}
	return TripleDesDecrypt(plainText, secretKey, ivAes)
}

func TripleDesDecryptByHex(cipherTextHex string, secretKey, ivAes []byte) ([]byte, error) {
	plainText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return []byte{}, err
	}
	return TripleDesDecrypt(plainText, secretKey, ivAes)
}
