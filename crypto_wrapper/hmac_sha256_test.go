package crypto_wrapper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHMAC(t *testing.T) {
	// Give a key and plain message
	key := []byte("Hello I am a 32 byte key........")
	plainMessage := []byte("Hello I am the plain message that will gets encrpyted")

	Convey("HMACSHA256 Should Always Produce The Same Result For Same Message And Same Key", t, func() {
		mac := HMACSHA256(key, plainMessage)
		tag := HMACSHA256(key, plainMessage)
		So(mac, ShouldResemble, tag)
	})
}
