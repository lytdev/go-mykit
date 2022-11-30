package gcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

// aesEcb  不建议使用ECB模式,推荐使用CBC模式
type aesEcb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *aesEcb {
	return &aesEcb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter aesEcb

func newECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		return
	}
	if len(dst) < len(src) {
		return
	}

	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter aesEcb

func newECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int {
	return x.blockSize
}

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		return
	}
	if len(dst) < len(src) {
		return
	}

	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func AesEcbEncrypt(plainText, secretKey []byte) ([]byte, error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	paddingText := PKCS5Padding(plainText, block.BlockSize())

	encrypted := make([]byte, len(paddingText))
	encrypter := newECBEncrypter(block)
	encrypter.CryptBlocks(encrypted, paddingText)

	return encrypted, nil
}

func AesEcbDecrypt(plainText, secretKey []byte) ([]byte, error) {
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return nil, ErrKeyLengthSixteen
	}
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	decrypter := newECBDecrypter(block)
	decrypted := make([]byte, len(plainText))
	decrypter.CryptBlocks(decrypted, plainText)

	return PKCS5UnPadding(decrypted), nil
}

func AesEcbEncryptBase64(plainText, key []byte) (string, error) {
	encrypted, err := AesEcbEncrypt(plainText, key)
	return base64.StdEncoding.EncodeToString(encrypted), err
}

func AesEcbEncryptHex(plainText, key []byte) (string, error) {
	encrypted, err := AesEcbEncrypt(plainText, key)
	return hex.EncodeToString(encrypted), err
}

func AesEcbDecryptByBase64(cipherTextBase64 string, key []byte) ([]byte, error) {
	plainText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return []byte{}, err
	}
	return AesEcbDecrypt(plainText, key)
}

func AesEcbDecryptByHex(cipherTextHex string, key []byte) ([]byte, error) {
	plainText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return []byte{}, err
	}
	return AesEcbDecrypt(plainText, key)
}
