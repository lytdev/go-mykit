package gcrypto

import (
	"errors"
)

const (
	// 密钥的长度

	KeyLength8  = 8
	KeyLength16 = 16
	KeyLength24 = 24
	KeyLength32 = 32
)

var (
	ErrCipherKey           = errors.New("the secret key is wrong and cannot be decrypted. Please check")
	ErrKeyLengthSixteen    = errors.New("a sixteen or twenty-four or thirty-two length secret key is required")
	ErrKeyLengthEight      = errors.New("a eight-length secret key is required")
	ErrKeyLengthTwentyFour = errors.New("a twenty-four-length secret key is required")
	ErrPaddingSize         = errors.New("padding size error please check the secret key or iv")
	ErrIvAes               = errors.New("a sixteen-length ivaes is required")
	ErrIvDes               = errors.New("a eight-length ivdes key is required")
	ErrRsaBits             = errors.New("bits 1024 or 2048")
	ErrAesSrcBlockSize     = errors.New("the length of src must be a multiple of the block size")
)

var (
	ErrImport                     = errors.New("ecies: failed to import key")
	ErrInvalidCurve               = errors.New("ecies: invalid elliptic curve")
	ErrInvalidParams              = errors.New("ecies: invalid ECIES parameters")
	ErrInvalidPublicKey           = errors.New("ecies: invalid public key")
	ErrSharedKeyIsPointAtInfinity = errors.New("ecies: shared key is point at infinity")
	ErrSharedKeyTooBig            = errors.New("ecies: shared key params are too big")
	ErrUnsupportedECIESParameters = errors.New("ecies: unsupported ECIES parameters")
	ErrKeyDataTooLong             = errors.New("ecies: can't supply requested key data")
	ErrSharedTooLong              = errors.New("ecies: shared secret is too long")
	ErrInvalidMessage             = errors.New("ecies: invalid message")
)
