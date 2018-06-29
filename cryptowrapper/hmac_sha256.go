package cryptowrapper

import (
	"crypto/hmac"
	"crypto/sha256"
)

// HMACSHA256 generates the HMAC for a message given a key.
func HMACSHA256(key []byte, messages ...[]byte) []byte {
	mac := hmac.New(sha256.New, key)
	for _, message := range messages {
		mac.Write(message)
	}
	messageMac := mac.Sum(nil)
	return messageMac
}
