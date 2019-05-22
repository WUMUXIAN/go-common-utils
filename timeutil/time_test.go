package timeutil

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeUtilTime(t *testing.T) {
	Convey("Test Time Utilities For Time", t, func() {
		Convey("ParseDataTimeISO8601 Should Work", func() {
			t, err := ParseDataTimeISO8601("2017-07-04 03:32:31")
			So(err, ShouldBeNil)
			So(t.Year(), ShouldEqual, 2017)
			So(t.Month(), ShouldEqual, 7)
			So(t.Day(), ShouldEqual, 4)
			So(t.Hour(), ShouldEqual, 3)
			So(t.Minute(), ShouldEqual, 32)
			So(t.Second(), ShouldEqual, 31)

			_, err = ParseDataTimeISO8601("12345")
			So(err, ShouldNotBeNil)
		})

		Convey("Parse Timestamp to Time", func() {
			// timestamp in seconds
			t := ParseTimeStamp(1499139151)
			So(t.UTC().Year(), ShouldEqual, 2017)
			So(t.UTC().Month(), ShouldEqual, 7)
			So(t.UTC().Day(), ShouldEqual, 4)
			So(t.UTC().Hour(), ShouldEqual, 3)
			So(t.UTC().Minute(), ShouldEqual, 32)
			So(t.UTC().Second(), ShouldEqual, 31)
			So(t.UTC().Unix(), ShouldEqual, 1499139151)
			So(t.UTC().UnixNano(), ShouldEqual, 1499139151000000000)

			// timestamp in milli seconds
			t = ParseTimeStamp(1499139151987)
			So(t.UTC().Year(), ShouldEqual, 2017)
			So(t.UTC().Month(), ShouldEqual, 7)
			So(t.UTC().Day(), ShouldEqual, 4)
			So(t.UTC().Hour(), ShouldEqual, 3)
			So(t.UTC().Minute(), ShouldEqual, 32)
			So(t.UTC().Second(), ShouldEqual, 31)
			So(t.UTC().Unix(), ShouldEqual, 1499139151)
			So(t.UTC().UnixNano(), ShouldEqual, 1499139151987000000)

			// timestamp in micro seconds
			t = ParseTimeStamp(1499139151987112)
			So(t.UTC().Year(), ShouldEqual, 2017)
			So(t.UTC().Month(), ShouldEqual, 7)
			So(t.UTC().Day(), ShouldEqual, 4)
			So(t.UTC().Hour(), ShouldEqual, 3)
			So(t.UTC().Minute(), ShouldEqual, 32)
			So(t.UTC().Second(), ShouldEqual, 31)
			So(t.UTC().Unix(), ShouldEqual, 1499139151)
			So(t.UTC().UnixNano(), ShouldEqual, 1499139151987112000)

			// timestamp in nano seconds
			t = ParseTimeStamp(1499139151987112565)
			So(t.UTC().Year(), ShouldEqual, 2017)
			So(t.UTC().Month(), ShouldEqual, 7)
			So(t.UTC().Day(), ShouldEqual, 4)
			So(t.UTC().Hour(), ShouldEqual, 3)
			So(t.UTC().Minute(), ShouldEqual, 32)
			So(t.UTC().Second(), ShouldEqual, 31)
			So(t.UTC().Unix(), ShouldEqual, 1499139151)
			So(t.UTC().UnixNano(), ShouldEqual, 1499139151987112565)
		})

		Convey("Parse MySQl DateTime String To Time", func() {
			// mysql datetime string without fractional seconds
			t, err := ParseDataTimeMySQL("2017-07-04 03:32:31")
			So(err, ShouldBeNil)
			So(t.Year(), ShouldEqual, 2017)
			So(t.Month(), ShouldEqual, 7)
			So(t.Day(), ShouldEqual, 4)
			So(t.Hour(), ShouldEqual, 3)
			So(t.Minute(), ShouldEqual, 32)
			So(t.Second(), ShouldEqual, 31)
			name, offset := t.Zone()
			So(name, ShouldEqual, "UTC")
			So(offset, ShouldEqual, 0)
			So(t.UTC().Unix(), ShouldEqual, 1499139151)
			So(t.UTC().UnixNano(), ShouldEqual, 1499139151000000000)

			// mysql datetime string with fractional seconds
			t, err = ParseDataTimeMySQL("2017-07-04 03:32:31.987")
			So(err, ShouldBeNil)
			So(t.Year(), ShouldEqual, 2017)
			So(t.Month(), ShouldEqual, 7)
			So(t.Day(), ShouldEqual, 4)
			So(t.Hour(), ShouldEqual, 3)
			So(t.Minute(), ShouldEqual, 32)
			So(t.Second(), ShouldEqual, 31)
			name, offset = t.Zone()
			So(name, ShouldEqual, "UTC")
			So(offset, ShouldEqual, 0)
			So(t.UTC().Unix(), ShouldEqual, 1499139151)
			So(t.UTC().UnixNano(), ShouldEqual, 1499139151987000000)

			// mysql datetime string with fractional seconds
			t, err = ParseDataTimeMySQL("2017-07-04 03:32:31.987112")
			So(err, ShouldBeNil)
			So(t.Year(), ShouldEqual, 2017)
			So(t.Month(), ShouldEqual, 7)
			So(t.Day(), ShouldEqual, 4)
			So(t.Hour(), ShouldEqual, 3)
			So(t.Minute(), ShouldEqual, 32)
			So(t.Second(), ShouldEqual, 31)
			name, offset = t.Zone()
			So(name, ShouldEqual, "UTC")
			So(offset, ShouldEqual, 0)
			So(t.UTC().Unix(), ShouldEqual, 1499139151)
			So(t.UTC().UnixNano(), ShouldEqual, 1499139151987112000)

			// invalid formats
			_, err = ParseDataTimeMySQL("2017-07-04")
			So(err, ShouldNotBeNil)
			_, err = ParseDataTimeMySQL("12345")
			So(err, ShouldNotBeNil)
		})
	})
}
