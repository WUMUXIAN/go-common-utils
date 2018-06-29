package cryptowrapper

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUUID(t *testing.T) {
	var re = regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)

	Convey("Gen UUID Should Have Correct Format And Should Be Unique", t, func() {
		uuid1 := GenUUID()
		uuid2 := GenUUID()
		So(uuid1, ShouldNotEqual, uuid2)
		So(re.MatchString(uuid1), ShouldBeTrue)
		So(re.MatchString(uuid2), ShouldBeTrue)
	})

	Convey("MD5UUIDFormat Should Give The Right Format", t, func() {
		uuid := MD5UUIDFormat(RandBytes(16))
		So(re.MatchString(uuid), ShouldBeTrue)
	})

	Convey("Gen Random BigInt Should Be Successful", t, func() {
		bigInt := RandBigInt(10)
		So(len(bigInt.Bytes()), ShouldBeGreaterThan, 0)
	})
}
