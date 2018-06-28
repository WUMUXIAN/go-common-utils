package crypto

import (
	"crypto/rand"
	"io"
	"math/big"

	"github.com/WUMUXIAN/go-common-utils/codec"
)

// RandBytes returns n bytes of cryptographically strong random bytes.
func RandBytes(n int) []byte {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic("Random source is broken!")
	}
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
func GenUUID() string {
	randBytes := RandBytes(16)
	return BytesToUUIDFormat(randBytes)
}

// BytesToUUIDFormat converts a 16 bytes to UUID format
func BytesToUUIDFormat(bytes []byte) string {
	return codec.ToByteArray(bytes[0:4]).Hex() + "-" +
		codec.ToByteArray(bytes[4:6]).Hex() + "-" +
		codec.ToByteArray(bytes[6:8]).Hex() + "-" +
		codec.ToByteArray(bytes[8:10]).Hex() + "-" +
		codec.ToByteArray(bytes[10:]).Hex()
}

// MD5UUID gets the MD5 hash of the input content and convert to UUID Format
func MD5UUIDFormat(input ...[]byte) string {
	return BytesToUUIDFormat(codec.GetHash(codec.MD5, input...).Bytes())
}
