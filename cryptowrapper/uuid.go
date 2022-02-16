package cryptowrapper

import (
	"crypto/rand"
	"io"
	"math/big"
	"strings"

	"github.com/google/uuid"

	"github.com/TectusDreamlab/go-common-utils/codec"
)

// RandBytes returns n bytes of cryptographically strong random bytes.
func RandBytes(n int) []byte {
	b := make([]byte, n)
	io.ReadFull(rand.Reader, b)
	return b
}

// RandBigInt generates and return a bigInt 'bits' bits in length
func RandBigInt(bits int) *big.Int {
	n := bits / 8
	if 0 != bits%8 {
		n++
	}
	b := RandBytes(n)
	r := big.NewInt(0).SetBytes(b)
	return r
}

// GenUUID generates UUID
func GenUUIDAlt() string {
	randBytes := RandBytes(16)
	return BytesToUUIDFormat(randBytes)
}

// GenUUID generates UUID
func GenUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

// BytesToUUIDFormat converts a 16 bytes to UUID format
func BytesToUUIDFormat(bytes []byte) string {
	return codec.ToByteArray(bytes[0:4]).Hex() + "-" +
		codec.ToByteArray(bytes[4:6]).Hex() + "-" +
		codec.ToByteArray(bytes[6:8]).Hex() + "-" +
		codec.ToByteArray(bytes[8:10]).Hex() + "-" +
		codec.ToByteArray(bytes[10:]).Hex()
}

// MD5UUIDFormat gets the MD5 hash of the input content and convert to UUID Format
func MD5UUIDFormat(input ...[]byte) string {
	return BytesToUUIDFormat(codec.GetHash(codec.MD5, input...).Bytes())
}

// Validates UUID
func ValidateUUID(UUID string, sanitize ...bool) (string, error) {
	u, err := uuid.Parse(UUID)
	if err != nil {
		return "", err
	}

	uStr := u.String()
	if len(sanitize) > 0 && sanitize[0] {
		uStr = strings.ToLower(uStr)
	}

	return uStr, err
}
