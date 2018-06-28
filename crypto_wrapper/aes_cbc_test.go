package crypto_wrapper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAESCBC(t *testing.T) {
	// Given a key and a plain message, encrpyt it using
	key := []byte("Hello I am a 32 byte key........Hello I am a 32 byte mac key....")
	plainMessage := []byte("我们")

	var encryptedMessage []byte
	var decryptedMessage []byte
	var err error

	Convey("Test Unpadding", t, func() {
		var unpadded []byte
		var padded []byte

		Convey("Unpad Empty Byte Array Should Return Empty Byte Array", func() {
			unpadded = unpad([]byte{})
			So(len(unpadded), ShouldEqual, 0)
		})

		Convey("Unpad Illegal Padded Array Should Return UnChanged Byte Array", func() {
			Convey("When The Last Byte Turn Into Integer Is Larger Than The Array Length", func() {
				padded = make([]byte, 10)
				padded[len(padded)-1] = byte(13)
				unpadded = unpad(padded)
				So(unpadded, ShouldResemble, padded)
			})

			Convey("When The Last Byte Turn Into Integer Is Larger Than BlockSize (16)", func() {
				padded = make([]byte, 23)
				padded[len(padded)-1] = byte(18)
				unpadded = unpad(padded)
				So(unpadded, ShouldResemble, padded)
			})

			Convey("When The Last Byte Turn Into Integer Is Zero", func() {
				padded = make([]byte, 23)
				padded[len(padded)-1] = byte(0)
				unpadded = unpad(padded)
				So(unpadded, ShouldResemble, padded)
			})

			Convey("When The Padding Format Is Wrong", func() {
				padded = make([]byte, 18)
				padded[len(padded)-1] = byte(3)
				unpadded = unpad(padded)
				So(unpadded, ShouldResemble, padded)
			})
		})
	})

	Convey("Test AESCBCEncrpyt", t, func() {
		encryptedMessage, err = AESCBCEncrypt(key[:32], plainMessage)
		Convey("AESCBCEncrpyt A Plain Message With Given Key Should Be Successful", func() {
			So(err, ShouldBeNil)
		})

		Convey("AESCBCEncrpyt A Plain Message With Wrong Key Size Should Not Be Successful", func() {
			_, err = AESCBCEncrypt(key[:33], plainMessage)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Test AESCBCDecrypt", t, func() {
		Convey("AESCBCDecrypt A Encrypted Message With Correct Key Should Be Successful", func() {
			decryptedMessage, err = AESCBCDecrypt(key[:32], encryptedMessage)
			So(err, ShouldBeNil)
			So(decryptedMessage, ShouldResemble, plainMessage)
			So(err, ShouldBeNil)
		})

		Convey("AESCBCDecrypt Wrong Message Should Not Be Successful", func() {
			_, err = AESCBCDecrypt(key[:32], RandBytes(42))
			So(err, ShouldNotBeNil)
			_, err = AESCBCDecrypt(key[:32], RandBytes(16))
			So(err, ShouldNotBeNil)
		})

		Convey("AESCBCDecrypt A Encrypted Message With Wrong Key Should Not Be Successful", func() {
			_, err = AESCBCDecrypt(key[:33], encryptedMessage)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Test AESCBCEncryptWithMAC And AESCBCDecryptWithMAC", t, func() {
		encryptedMessage, err = AESCBCEncryptWithMAC(key, plainMessage)
		Convey("AESCBCEncryptWithMAC A Plain Message With Given Key Should Be Successful", func() {
			So(err, ShouldBeNil)
		})

		// decrpyt it with the right key
		Convey("AESCBCDecryptWithMAC The Encrypted Message With The Correct Key Should Be Successful", func() {
			decryptedMessage, err = AESCBCDecryptWithMAC(key, encryptedMessage)
			So(err, ShouldBeNil)
			So(decryptedMessage, ShouldResemble, plainMessage)
		})

		// decrypt it with the wrong key
		Convey("AESCBCDecryptWithMAC The Encrypted Message With The Wrong Key Should Not Be Successful", func() {
			decryptedMessage, err = AESCBCDecryptWithMAC(RandBytes(32), encryptedMessage)
			So(err, ShouldNotBeNil)
		})
	})
}
