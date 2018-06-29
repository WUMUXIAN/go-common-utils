package cryptowrapper

import (
	"crypto/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestECDH(t *testing.T) {
	Convey("Test ECDH Encryption", t, func() {
		clientPrivateKey, clientPublicKey, err := GenerateECDHKeyPair(rand.Reader)
		Convey("Generating ECDH Key Pair For Client Should Be OK", func() {
			So(err, ShouldBeNil)
		})
		serverPrivateKey, serverPublicKey, err := GenerateECDHKeyPair(rand.Reader)
		Convey("Generating ECDH Key Pair For Server Should Be OK", func() {
			So(err, ShouldBeNil)
		})

		Convey("Marshal And UnMarshal Public Key Should Be Successful", func() {
			clientPublicKeyBytes := MarshalECDHPublicKey(clientPublicKey)
			So(clientPublicKeyBytes, ShouldNotBeEmpty)
			clientPublicKeyNew, ok := UnmarshalECDHPublicKey(clientPublicKeyBytes)
			So(ok, ShouldBeTrue)
			So(clientPublicKeyNew, ShouldResemble, clientPublicKey)
		})

		Convey("UnMarshal Wrong Public Key Should Be Failed", func() {
			_, ok := UnmarshalECDHPublicKey(RandBytes(16))
			So(ok, ShouldBeFalse)
		})

		Convey("Client Private + Server Public And Client Public + Server Private Should Generate Same Shared Secret", func() {
			secret1, err := GenerateECDHSharedSecret(clientPrivateKey, serverPublicKey)
			So(err, ShouldBeNil)
			secret2, err := GenerateECDHSharedSecret(serverPrivateKey, clientPublicKey)
			So(err, ShouldBeNil)
			So(secret1, ShouldResemble, secret2)
		})

		Convey("Mismatched Public + Private Pair Should Not Generate Same Shared Secret", func() {
			secret1, err := GenerateECDHSharedSecret(clientPrivateKey, clientPublicKey)
			So(err, ShouldBeNil)
			secret2, err := GenerateECDHSharedSecret(serverPrivateKey, serverPublicKey)
			So(err, ShouldBeNil)
			So(secret1, ShouldNotResemble, secret2)
		})
	})
}
