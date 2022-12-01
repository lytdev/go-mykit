package gcrypto

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"github.com/lytdev/go-mykit/gnum"
)

// DesCbcEncrypt DES加密
func DesCbcEncrypt(plainText, secretKey, ivDes []byte) (cipherText []byte, err error) {
	if len(secretKey) != KeyLength8 {
		return nil, ErrKeyLengthEight
	}
	block, err := des.NewCipher(secretKey)
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

	cipherText = make([]byte, ptLen)
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}

// DesCbcDecrypt DES解密
func DesCbcDecrypt(cipherText, secretKey, ivDes []byte) (plainText []byte, err error) {
	if len(secretKey) != 8 {
		return nil, ErrKeyLengthEight
	}
	block, err := des.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	ctLen := len(cipherText)
	bLen := block.BlockSize()
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

	plainText = make([]byte, ctLen)
	blockMode.CryptBlocks(plainText, cipherText)
	return PKCS5UnPadding(plainText, bLen)
}

func DesCbcEncryptBase64(plainText, secretKey, ivAes []byte) (cipherTextBase64 string, err error) {
	encrypted, err := DesCbcEncrypt(plainText, secretKey, ivAes)
	return base64.StdEncoding.EncodeToString(encrypted), err
}

func DesCbcEncryptHex(plainText, secretKey, ivAes []byte) (cipherTextHex string, err error) {
	encrypted, err := DesCbcEncrypt(plainText, secretKey, ivAes)
	return hex.EncodeToString(encrypted), err
}

func DesCbcDecryptByBase64(cipherTextBase64 string, secretKey, ivAes []byte) (plainText []byte, err error) {
	plainTextBytes, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return []byte{}, err
	}
	return DesCbcDecrypt(plainTextBytes, secretKey, ivAes)
}

func DesCbcDecryptByHex(cipherTextHex string, secretKey, ivAes []byte) (plainText []byte, err error) {
	plainTextBytes, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return []byte{}, err
	}
	return DesCbcDecrypt(plainTextBytes, secretKey, ivAes)
}
