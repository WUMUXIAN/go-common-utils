package cryptowrapper

import (
	"crypto"
	"io"

	"golang.org/x/crypto/curve25519"
)

// Curve25519 is a state-of-the-art Diffie-Hellman function suitable for a wide variety of applications.
// Given a user's 32-byte secret key, Curve25519 computes the user's 32-byte public key.
// Given the user's 32-byte secret key and another user's 32-byte public key, Curve25519 computes a 32-byte secret shared by the two users.
// This secret can then be used to authenticate and encrypt messages between the two users.

// GenerateECDHKeyPair generates the ECDH key pair
func GenerateECDHKeyPair(rand io.Reader) (crypto.PrivateKey, crypto.PublicKey, error) {
	var pub, priv [32]byte
	var err error

	_, err = io.ReadFull(rand, priv[:])
	if err != nil {
		return nil, nil, err
	}
	curve25519.ScalarBaseMult(&pub, &priv)
	return &priv, &pub, nil
}

// MarshalECDHPublicKey marshals the public key into byte array
func MarshalECDHPublicKey(p crypto.PublicKey) []byte {
	pub := p.(*[32]byte)
	return pub[:]
}

// UnmarshalECDHPublicKey unmarshals byte array into public key
func UnmarshalECDHPublicKey(data []byte) (crypto.PublicKey, bool) {
	var pub [32]byte
	if len(data) != 32 {
		return nil, false
	}

	copy(pub[:], data)
	return &pub, true
}

// GenerateECDHSharedSecret calculates a common secret using the private key and public key
func GenerateECDHSharedSecret(privKey crypto.PrivateKey, pubKey crypto.PublicKey) ([]byte, error) {
	var priv, pub, secret *[32]byte

	priv = privKey.(*[32]byte)
	pub = pubKey.(*[32]byte)
	secret = new([32]byte)

	curve25519.ScalarMult(secret, priv, pub)
	return secret[:], nil
}
