package crypto

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPBKDF2_HMAC_SHA256(t *testing.T) {
	Convey("TestPBKDF2_HMAC_SHA256 Should Be Able To Derive Strong Password", t, func() {
		strongPassword := PBKDF2HMACSHA256([]byte("weakpassword"), RandBytes(16), 100000, 64)
		So(len(strongPassword), ShouldEqual, 64)
		So(strongPassword, ShouldNotResemble, []byte("weakpassword"))
	})
}
