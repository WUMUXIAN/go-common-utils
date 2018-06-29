package cryptowrapper

import (
	"crypto/sha256"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
}
