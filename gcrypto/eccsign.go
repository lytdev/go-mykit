package gcrypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"math/big"
)

// eccSign ECC签名
func eccSign(msg []byte, priKey []byte) (rSign []byte, sSign []byte, err error) {
	privateKey, err := x509.ParseECPrivateKey(priKey)
	if err != nil {
		return nil, nil, err
	}
	resultHash := Sha256(msg)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, resultHash)
	if err != nil {
		return nil, nil, err
	}

	rText, err := r.MarshalText()
	if err != nil {
		return nil, nil, err
	}
	sText, err := s.MarshalText()
	if err != nil {
		return nil, nil, err
	}
	return rText, sText, nil
}

// eccVerifySign ECC验证签名
func eccVerifySign(msg []byte, pubKey []byte, rText, sText []byte) bool {
	publicKeyInterface, _ := x509.ParsePKIXPublicKey(pubKey)
	publicKey := publicKeyInterface.(*ecdsa.PublicKey)
	resultHash := Sha256(msg)

	var r, s big.Int
	err := r.UnmarshalText(rText)
	if err != nil {
		return false
	}
	err = s.UnmarshalText(sText)
	if err != nil {
		return false
	}
	result := ecdsa.Verify(publicKey, resultHash, &r, &s)
	return result
}

// EccSignBase64 ECC签名转base64
func EccSignBase64(msg []byte, base64PriKey string) (base64rSign, base64sSign string, err error) {
	priBytes, err := base64.StdEncoding.DecodeString(base64PriKey)
	if err != nil {
		return "", "", err
	}
	rSign, sSign, err := eccSign(msg, priBytes)
	if err != nil {
		return "", "", err
	}
	return base64.StdEncoding.EncodeToString(rSign), base64.StdEncoding.EncodeToString(sSign), nil
}

// EccVerifySignBase64 ECC验证base64格式的签名
func EccVerifySignBase64(msg []byte, base64rSign, base64sSign, base64PubKey string) bool {
	rSignBytes, err := base64.StdEncoding.DecodeString(base64rSign)
	if err != nil {
		return false
	}
	sSignBytes, err := base64.StdEncoding.DecodeString(base64sSign)
	if err != nil {
		return false
	}
	pubBytes, err := base64.StdEncoding.DecodeString(base64PubKey)
	if err != nil {
		return false
	}
	return eccVerifySign(msg, pubBytes, rSignBytes, sSignBytes)
}

// EccSignHex ECC签名转Hex
func EccSignHex(msg []byte, hexPriKey string) (hexrSign, hexsSign string, err error) {
	priBytes, err := hex.DecodeString(hexPriKey)
	if err != nil {
		return "", "", err
	}
	rSign, sSign, err := eccSign(msg, priBytes)
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(rSign), hex.EncodeToString(sSign), nil
}

// EccVerifySignHex ECC验证hex格式的签名
func EccVerifySignHex(msg []byte, hexrSign, hexsSign, hexPubKey string) bool {
	rSignBytes, err := hex.DecodeString(hexrSign)
	if err != nil {
		return false
	}
	sSignBytes, err := hex.DecodeString(hexsSign)
	if err != nil {
		return false
	}
	pubBytes, err := hex.DecodeString(hexPubKey)
	if err != nil {
		return false
	}
	return eccVerifySign(msg, pubBytes, rSignBytes, sSignBytes)
}
