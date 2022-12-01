package gcrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
)

/*
Asymmetric encryption requires the generation of a pair of keys rather than a key, so before encryption here you need to get a pair of keys, public and private, respectively
Generate the public and private keys all at once
Encryption: plaintext to the power E Mod N to output ciphertext
Decryption: ciphertext to the power D Mod N outputs plaintext
Encryption operations take a long time? Encryption is faster
The data is encrypted and cannot be easily decrypted
*/

type RsaKey struct {
	PrivateKey string
	PublicKey  string
}

const (
	keyLen1024 = 1024
	keyLen2048 = 2048
)

// GenerateRsaKey 生成公钥和私钥
func GenerateRsaKey(bits int) (*rsa.PrivateKey, error) {
	if bits != keyLen1024 && bits != keyLen2048 {
		return nil, ErrRsaBits
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// GenerateRsaKeyHex 生成公钥和私钥的hex格式
func GenerateRsaKeyHex(bits int) (RsaKey, error) {
	privateKey, err := GenerateRsaKey(bits)
	if err != nil {
		return RsaKey{}, err
	}
	return RsaKey{
		PrivateKey: hex.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  hex.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}

// GenerateRsaKeyBase64 生成公钥和私钥的base64格式
func GenerateRsaKeyBase64(bits int) (RsaKey, error) {
	privateKey, err := GenerateRsaKey(bits)
	if err != nil {
		return RsaKey{}, err
	}
	return RsaKey{
		PrivateKey: base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}
