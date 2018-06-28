package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

// PBKDF2 (Password-Based Key Derivation Function 2) is part of RSA Laboratories' Public-Key Cryptography Standards (PKCS) series,
// specifically PKCS #5 v2.0, also published as Internet Engineering Task Force's RFC 2898.
// It replaces an earlier key derivation function, PBKDF1, which could only produce derived keys up to 160 bits long.

// PBKDF2HMACSHA256 derives a key for a given password
func PBKDF2HMACSHA256(password, salt []byte, iter, keyLen int) []byte {
	return pbkdf2.Key(password, salt, iter, keyLen, sha256.New)
}
