package crypto_wrapper

import (
	"hash"
	"io"

	"golang.org/x/crypto/hkdf"
)

// HKDF is a cryptographic key derivation function (KDF) with the goal of expanding limited
// input keying material into one or more cryptographically strong secret keys.

// DeriveKey derives two cryptographically strong secret keys from the weak key.
// salt and info can be nil.
// salt is recommended to be hash-length sized random
// info is recommended to independent from the weak key.
// keySize is the desired size for the derived key, its up to yourself to split it as multiple keys for usage.
func DeriveKey(hash func() hash.Hash, weakKey []byte, salt []byte, info []byte, keySize int) ([]byte, error) {

	hkdf := hkdf.New(hash, weakKey, salt, info)

	// Generate the required key
	key := make([]byte, keySize)
	n, err := io.ReadFull(hkdf, key)
	if n != len(key) || err != nil {
		return key, err
	}
	return key, nil
}
