package gcrypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
)

func rsaSign(msg, priKey []byte) (sign []byte, err error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(priKey)
	hashed := Sha256(msg)
	sign, err = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

func rsaVerifySign(msg []byte, sign []byte, pubKey []byte) bool {
	publicKey, err := x509.ParsePKCS1PublicKey(pubKey)
	if err != nil {
		return false
	}
	hashed := Sha256(msg)
	result := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, sign)
	return result == nil
}

func RsaSignBase64(msg []byte, base64PriKey string) (base64Sign string, err error) {
	priBytes, err := base64.StdEncoding.DecodeString(base64PriKey)
	if err != nil {
		return "", err
	}
	sign, err := rsaSign(msg, priBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func RsaVerifySignBase64(msg []byte, base64Sign, base64PubKey string) bool {
	signBytes, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return false
	}
	pubBytes, err := base64.StdEncoding.DecodeString(base64PubKey)
	if err != nil {
		return false
	}
	return rsaVerifySign(msg, signBytes, pubBytes)
}

func RsaSignHex(msg []byte, hexPriKey string) (hexSign string, err error) {
	priBytes, err := hex.DecodeString(hexPriKey)
	if err != nil {
		return "", err
	}
	sign, err := rsaSign(msg, priBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sign), nil
}

func RsaVerifySignHex(msg []byte, hexSign, hexPubKey string) bool {
	signBytes, err := hex.DecodeString(hexSign)
	if err != nil {
		return false
	}
	pubBytes, err := hex.DecodeString(hexPubKey)
	if err != nil {
		return false
	}
	return rsaVerifySign(msg, signBytes, pubBytes)
}
