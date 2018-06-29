package cryptowrapper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// AESCBCEncrypt performs a AES encryption in CBC mode
func AESCBCEncrypt(key []byte, input []byte) ([]byte, error) {
	paddedInput := pad(input)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	output := make([]byte, aes.BlockSize+len(paddedInput))
	iv := output[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(output[aes.BlockSize:], paddedInput)

	return output, nil
}

// AESCBCDecrypt performs a AES decryption in CBC mode
func AESCBCDecrypt(key []byte, input []byte) ([]byte, error) {
	if len(input)%aes.BlockSize != 0 {
		return nil, errors.New("Input data size is wrong")
	}

	// It will be at least 2 aes.BlockSize
	if len(input) < (2 * aes.BlockSize) {
		return nil, errors.New("Input data size is wrong")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := input[:aes.BlockSize]
	input = input[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(input, input)

	output := unpad(input)

	return output, err
}

// AESCBCEncryptWithMAC performs a AES encryption in CBC mode with a MAC authentication
func AESCBCEncryptWithMAC(key []byte, input []byte) ([]byte, error) {
	aesKey := key[:32]
	hmacKey := key[32:]

	encrypted, err := AESCBCEncrypt(aesKey, input)

	if err != nil {
		return nil, err
	}

	mac := HMACSHA256(hmacKey, encrypted)
	encrypted = append(encrypted, mac...)
	return encrypted, nil
}

// AESCBCDecryptWithMAC performs a AES dencryption in CBC mode with a MAC authentication
func AESCBCDecryptWithMAC(key []byte, input []byte) ([]byte, error) {
	aesKey := key[:32]
	hmacKey := key[32:]

	mac := input[len(input)-32:]
	input = input[:len(input)-32]

	if !bytes.Equal(HMACSHA256(hmacKey, input), mac) {
		return nil, errors.New("Signature failed")
	}

	// Now the verification passes, decrpyt the message.
	return AESCBCDecrypt(aesKey, input)
}

func pad(input []byte) []byte {
	padding := aes.BlockSize - (len(input) % aes.BlockSize)
	padded := make([]byte, len(input))
	copy(padded, input)
	for i := 0; i < padding; i++ {
		padded = append(padded, byte(padding))
	}
	return padded
}

func unpad(input []byte) []byte {
	if len(input) == 0 {
		return input
	}

	padding := input[len(input)-1]

	if int(padding) > len(input) || int(padding) > aes.BlockSize {
		return input
	} else if int(padding) == 0 {
		return input
	}

	for i := len(input) - 1; i > len(input)-int(padding)-1; i-- {
		if input[i] != padding {
			return input
		}
	}

	return input[:len(input)-int(padding)]
}
