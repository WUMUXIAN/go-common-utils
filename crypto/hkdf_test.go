package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"testing"

	"tecgit01.tectusdreamlab.com/TDS/common-utils-backend/convertor"

	. "github.com/smartystreets/goconvey/convey"
)

func testOne(t *testing.T) {
	weakKey := []byte{0x00, 0x01, 0x02, 0x03}
	hash := sha256.New
	salt := make([]byte, hash().Size())
	n, err := io.ReadFull(rand.Reader, salt)
	if n != len(salt) || err != nil {
		fmt.Println("error:", err)
		return
	}

	// Derive a key with length of 64
	derivedKeys, err := DeriveKey(hash, weakKey, salt, nil, 64)
	if err == nil {
		// Split the key into 2 parts, first 32 byte as an AES key and later 32 bytes as a HAMC key.
		fmt.Println("AES Key:", convertor.BytesToBase64(derivedKeys[:32]))
		fmt.Println("HAMC Key:", convertor.BytesToBase64(derivedKeys[32:]))
	}
}

func testTwo(t *testing.T) {
	weakKey := []byte{0x00, 0x01, 0x02, 0x03}
	hash := sha256.New
	// Derive a key with length of 80
	derivedKeys, err := DeriveKey(hash, weakKey, nil, nil, 80)
	if err == nil {
		// Split the key into 2 parts, first 32 byte as an AES key, next 32 bytes as a HAMC key and final 16 bytes as a VI
		fmt.Println("AES Key:", convertor.BytesToBase64(derivedKeys[:32]))
		fmt.Println("HAMC Key:", convertor.BytesToBase64(derivedKeys[32:64]))
		fmt.Println("VI:", convertor.BytesToBase64(derivedKeys[64:]))
	}
}

func testThree(t *testing.T) {
	weakKey := []byte{0x00, 0x01, 0x02, 0x03}
	hash := sha256.New
	// Derive 100 keys
	derivedKeys := make([][]byte, 100)
	for i := 0; i < 100; i++ {
		key, _ := DeriveKey(hash, weakKey, nil, nil, 32)
		derivedKeys = append(derivedKeys, key)
	}

	for i := 0; i < 99; i++ {
		for j := i + 1; j < 100; j++ {
			if !bytes.Equal(derivedKeys[i], derivedKeys[j]) {
				t.Fatal("HKDF function doesn't work")
			}
		}
	}
}

func TestHDKF(t *testing.T) {
	Convey("Test HKDF (HMAC-Based Key Derivation Function)", t, func() {
		weakKey := []byte{0x00, 0x01, 0x02, 0x03}
		hash := sha256.New
		salt := make([]byte, hash().Size())

		Convey("Derive A 4 Bytes Weak Key To A 64 Bytes Strong Key Should Be Successful", func() {
			derivedKeys, err := DeriveKey(hash, weakKey, salt, nil, 64)
			So(err, ShouldBeNil)
			So(len(derivedKeys), ShouldEqual, 64)
		})

		Convey("With The Same Key, Salt And Info, Derived Key Should Be The Same", func() {
			derivedKeys1, _ := DeriveKey(hash, weakKey, salt, nil, 90)
			derivedKeys2, _ := DeriveKey(hash, weakKey, salt, nil, 90)
			So(derivedKeys1, ShouldResemble, derivedKeys2)
		})
	})

	testTwo(t)
	testThree(t)
}
