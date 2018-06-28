package crypto_wrapper

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"errors"
)

// GenerateRSAKey generates the RSA private key
func GenerateRSAKey(bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

// RSAEncrypt encrypts the given message with the public key
func RSAEncrypt(pub *rsa.PublicKey, msg []byte, label []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, msg, label)
}

// RSADecrypt decrypts the given cipher text with the private key
func RSADecrypt(priv *rsa.PrivateKey, ciphertext []byte, label []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, label)
}

// RSASign signs the given message with the private key.
func RSASign(priv *rsa.PrivateKey, message []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(message)
	hashed := hash.Sum(nil)
	pssOpt := &rsa.PSSOptions{
		SaltLength: 32,
		Hash:       crypto.SHA256,
	}
	return rsa.SignPSS(rand.Reader, priv, crypto.SHA256, hashed, pssOpt)
}

// RSAVerify verifies the given message with the public key
func RSAVerify(pub *rsa.PublicKey, message []byte, signiture []byte) error {
	hash := sha256.New()
	hash.Write(message)
	hashed := hash.Sum(nil)
	pssOpt := &rsa.PSSOptions{
		SaltLength: 32,
		Hash:       crypto.SHA256,
	}
	return rsa.VerifyPSS(pub, crypto.SHA256, hashed, signiture, pssOpt)
}

// MarshalPublicKey marshals public key into ANS.1 DER bytes (PKSC#1 scheme)
func MarshalPublicKey(pub *rsa.PublicKey) ([]byte, error) {
	return asn1.Marshal(*pub)
}

// MarshalPrivateKey marshals private key into ANS.1 DER bytes (PKSC#1 scheme)
func MarshalPrivateKey(priv *rsa.PrivateKey) ([]byte, error) {
	return x509.MarshalPKCS1PrivateKey(priv), nil
}

// UnMarshalPublicKey unmarshals ANS.1 DER bytes back to public key
func UnMarshalPublicKey(der []byte) (pubKey *rsa.PublicKey, err error) {
	pubKey = &rsa.PublicKey{}
	if rest, err := asn1.Unmarshal(der, pubKey); err == nil {
		if len(rest) < 1 && err == nil {
			return pubKey, nil
		}
	}
	return nil, errors.New("wrong public key bytes")
}

// UnMarshalPrivateKey unmarshals ANS.1 DER bytes back to private key
func UnMarshalPrivateKey(der []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(der)
}
