package crypto_wrapper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRSA(t *testing.T) {
	Convey("Test RSA Encryption Algorithm", t, func() {
		priv, err := GenerateRSAKey(2048)
		priv1, _ := GenerateRSAKey(2048)

		Convey("Generate A RSA Key Pair Should Be Successful", func() {
			So(err, ShouldBeNil)
			So(len(priv.PublicKey.N.Bytes()), ShouldBeLessThanOrEqualTo, 256)
			So(priv.PublicKey.E, ShouldEqual, 65537)
			So(len(priv.D.Bytes()), ShouldBeLessThanOrEqualTo, 256)
			So(len(priv.Primes[0].Bytes()), ShouldBeLessThanOrEqualTo, 128)
			So(len(priv.Primes[0].Bytes()), ShouldBeLessThanOrEqualTo, 128)
			So(len(priv.Precomputed.Dp.Bytes()), ShouldBeLessThanOrEqualTo, 128)
			So(len(priv.Precomputed.Dq.Bytes()), ShouldBeLessThanOrEqualTo, 128)
			So(len(priv.Precomputed.Qinv.Bytes()), ShouldBeLessThanOrEqualTo, 128)
		})

		Convey("MarshalPrivateKey And UnMarshalPrivateKey On Valid Private Key Should Be Successful", func() {
			privBytes, err := MarshalPrivateKey(priv)
			So(err, ShouldBeNil)
			So(len(privBytes), ShouldBeGreaterThan, 0)
			priv, err := UnMarshalPrivateKey(privBytes)
			So(err, ShouldBeNil)
			So(priv, ShouldNotBeNil)
		})

		Convey("MarshalPublicKey And UnMarshalPublicKey On Valid Public Key Should Be Successful", func() {
			pubBytes, err := MarshalPublicKey(&priv.PublicKey)
			So(err, ShouldBeNil)
			So(len(pubBytes), ShouldBeGreaterThan, 0)
			pub, err := UnMarshalPublicKey(pubBytes)
			So(err, ShouldBeNil)
			So(priv.PublicKey.N.Bytes(), ShouldResemble, pub.N.Bytes())
		})

		Convey("UnMarshalPrivateKey On InValid Private Key Should Not Be Successful", func() {
			priv, err := UnMarshalPrivateKey(RandBytes(456))
			So(err, ShouldNotBeNil)
			So(priv, ShouldBeNil)
		})

		Convey("UnMarshalPublicKey On InValid Public Key Should Not Be Successful", func() {
			pub, err := UnMarshalPublicKey(RandBytes(256))
			So(err, ShouldNotBeNil)
			So(pub, ShouldBeNil)
		})

		Convey("Sign And Verify Should Work", func() {
			msg := []byte("Test message for RSA")

			signature1, err1 := RSASign(priv, msg)
			signature2, err2 := RSASign(priv, msg)
			Convey("Sign Should Work And Every Signature Should Be Different", func() {
				So(err1, ShouldBeNil)
				So(err2, ShouldBeNil)
				So(signature1, ShouldNotResemble, signature2)
			})

			Convey("Verify Should Work For All Signatures", func() {
				err := RSAVerify(&priv.PublicKey, msg, signature1)
				So(err, ShouldBeNil)
				err = RSAVerify(&priv.PublicKey, msg, signature2)
				So(err, ShouldBeNil)
			})

			Convey("Verify Should Fail If Key Pair Does Not Match", func() {
				err := RSAVerify(&priv1.PublicKey, msg, signature1)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Encrypt And Decrypte Should Work", func() {
			msg := []byte("Test message for RSA")

			encryptedMessage1, err1 := RSAEncrypt(&priv.PublicKey, msg, []byte("label"))
			encryptedMessage2, err2 := RSAEncrypt(&priv.PublicKey, msg, []byte("label"))
			encryptedMessage3, err3 := RSAEncrypt(&priv.PublicKey, msg, []byte(""))

			Convey("Encryption Should Work And Every Encrypted Message Should Be Different", func() {
				So(err1, ShouldBeNil)
				So(err2, ShouldBeNil)
				So(err3, ShouldBeNil)
				So(encryptedMessage1, ShouldNotResemble, encryptedMessage2)
				So(encryptedMessage2, ShouldNotResemble, encryptedMessage3)
			})

			Convey("Decryption Should Work And Every Encrypted Message Should Be Decrypted If Labels Are Same", func() {
				decryptedMessage, err := RSADecrypt(priv, encryptedMessage1, []byte("label"))
				So(err, ShouldBeNil)
				So(decryptedMessage, ShouldResemble, msg)
				decryptedMessage, err = RSADecrypt(priv, encryptedMessage2, []byte("label"))
				So(err, ShouldBeNil)
				So(decryptedMessage, ShouldResemble, msg)
				decryptedMessage, err = RSADecrypt(priv, encryptedMessage3, []byte(""))
				So(err, ShouldBeNil)
				So(decryptedMessage, ShouldResemble, msg)
				decryptedMessage, err = RSADecrypt(priv, encryptedMessage2, []byte(""))
				So(err, ShouldNotBeNil)
				So(decryptedMessage, ShouldNotResemble, msg)
			})

			Convey("Decryption Should Fail If Key Pair Does Not Match", func() {
				decryptedMessage, err := RSADecrypt(priv1, encryptedMessage1, []byte("label"))
				So(err, ShouldNotBeNil)
				So(decryptedMessage, ShouldNotResemble, msg)
			})
		})
	})
}
