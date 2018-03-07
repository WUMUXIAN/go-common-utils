package codec

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"math/big"
)

// HashType defines the supported hash type, currently support md5, sha1, sha256 and sha512
type HashType int

// Enum the HashType
const (
	MD5 HashType = iota
	SHA1
	SHA256
	SHA512
)

// ByteArray defines a byte arry
type ByteArray []byte

// Bytes gets the []byte
func (o ByteArray) Bytes() []byte {
	return []byte(o)
}

// Base64 returns the base64 encoded string of the byte array
func (o ByteArray) Base64() string {
	return base64.StdEncoding.EncodeToString(o.Bytes())
}

// Base64URL returns the base64 URL encoded string of the byte array
func (o ByteArray) Base64URL() string {
	return base64.URLEncoding.EncodeToString(o.Bytes())
}

// Base32 returns the base32 encoded string of the byte array
func (o ByteArray) Base32() string {
	return base32.StdEncoding.EncodeToString(o.Bytes())
}

// Hex returns the hex encoded string of the byte array
func (o ByteArray) Hex() string {
	return hex.EncodeToString(o.Bytes())
}

// BigInt returns big integer the byte array represents
func (o ByteArray) BigInt() *big.Int {
	i := big.NewInt(0)
	return i.SetBytes(o.Bytes())
}

// ToByteArray return the combined byte array for the input byte arrays.
func ToByteArray(input ...[]byte) ByteArray {
	sum := make([]byte, 0)
	for _, i := range input {
		sum = append(sum, i...)
	}
	return ByteArray(sum)
}

// GetHash returns the hash of the input given the hash type.
func GetHash(hashType HashType, input ...[]byte) ByteArray {
	var h hash.Hash
	switch hashType {
	case MD5:
		h = md5.New()
	case SHA1:
		h = sha1.New()
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	}
	for _, i := range input {
		h.Write(i)
	}
	return h.Sum(nil)
}

// HexToBytes converts hex string to byte array
func HexToBytes(input string) ([]byte, error) {
	return hex.DecodeString(input)
}

// Base64ToBytes converts base64 string to byte array
func Base64ToBytes(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}

// Base64URLToBytes converts base64URL string to byte array
func Base64URLToBytes(input string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(input)
}

// Base32ToBytes converts base32 string to byte array
func Base32ToBytes(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
