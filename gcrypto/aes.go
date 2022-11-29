package gcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/lytdev/go-mykit/gstr"
	"io"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

// key不能泄露 16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
// 秘钥长度验证
func validKey(key []byte) bool {
	k := len(key)
	switch k {
	default:
		return false
	case 16, 24, 32:
		return true
	}
}

// =================== CBC ======================

// AesEncryptCBC aes加密
func AesEncryptCBC(key, src []byte) (data []byte, err error) {
	if !validKey(key) {
		return nil, errors.New("密钥长度错误")
	}
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	} else if len(src) == 0 {
		return nil, errors.New("src is empty")
	}
	//判断加密快的大小,填充
	plaintext, err := pkcs7Padding(src, block.BlockSize())
	if err != nil {
		return nil, err
	}
	//初始化加密数据接收切片
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	//使用cbc加密模式
	bm := cipher.NewCBCEncrypter(block, iv)
	//执行加密
	bm.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// AesDecryptCBC aes解密
func AesDecryptCBC(key, src []byte) (data []byte, err error) {
	if !validKey(key) {
		return nil, errors.New("加密的密钥长度错误")
	}
	if len(src) < aes.BlockSize {
		return nil, errors.New("数据错误")
	}
	iv := src[:aes.BlockSize]
	ciphertext := src[aes.BlockSize:]
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//使用cbc解密
	bm := cipher.NewCBCDecrypter(block, iv)
	//执行解密
	bm.CryptBlocks(ciphertext, ciphertext)
	//去除填充
	ciphertext, err = pkcs7UnPadding(ciphertext, aes.BlockSize)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// =================== ECB ======================

// AesEncryptECB aes加密
func AesEncryptECB(key, origData []byte) (encrypted []byte) {
	block, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, block.BlockSize(); bs <= len(origData); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

// AesDecryptECB aes解密
func AesDecryptECB(key, encrypted []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, block.BlockSize(); bs < len(encrypted); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// =================== CFB ======================

// AesEncryptCFB aes加密
func AesEncryptCFB(key, origData []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}

// AesDecryptCFB aes解密
func AesDecryptCFB(key, encrypted []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}

//////////////////默认使用的加密和解密////////////////////////////

// EncryptDataByAes Aes加密后base64
func EncryptDataByAes(key string, data []byte) (string, error) {
	res, err := AesEncryptCBC(gstr.StrToByteArr(key), data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

// DecryptDataByAes base64解码后 Aes 解密
func DecryptDataByAes(key, data string) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return AesDecryptCBC(gstr.StrToByteArr(key), dataByte)
}
