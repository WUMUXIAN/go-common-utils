package timeutil

import (
	"testing"
	"time"

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

		Convey("Get Current Time In UTC", func() {
			t := NowUTC()
			compareT := time.Now().UTC()
			So(t.Year(), ShouldEqual, compareT.Year())
			So(t.Month(), ShouldEqual, compareT.Month())
			So(t.Day(), ShouldEqual, compareT.Day())
			So(t.Hour(), ShouldEqual, compareT.Hour())
			So(t.Minute(), ShouldEqual, compareT.Minute())
			So(t.Second(), ShouldEqual, compareT.Second())
			name, offset := t.Zone()
			So(name, ShouldEqual, "UTC")
			So(offset, ShouldEqual, 0)
			So(t.Unix(), ShouldEqual, compareT.Unix())
			So(t.UnixNano(), ShouldEqual, compareT.UnixNano())
		})

		Convey("Get Current Time In UTC - Add Duration", func() {
			t := NowUTC(time.Second * time.Duration(1))
			compareT := time.Now().UTC()
			So(t.Year(), ShouldEqual, compareT.Year())
			So(t.Month(), ShouldEqual, compareT.Month())
			So(t.Day(), ShouldEqual, compareT.Day())
			So(t.Hour(), ShouldEqual, compareT.Hour())
			So(t.Minute(), ShouldEqual, compareT.Minute())
			So(t.Second(), ShouldEqual, compareT.Second()+1)
			name, offset := t.Zone()
			So(name, ShouldEqual, "UTC")
			So(offset, ShouldEqual, 0)
			So(t.Unix(), ShouldEqual, compareT.Unix()+1)
			So(t.UnixNano(), ShouldEqual, compareT.UnixNano()+1*1000000000)
		})

		Convey("Check The MySQL Date Time Formatting", func() {
			t := time.Unix(0, 1574309440240*int64(time.Millisecond))
			So(t.UTC().Format(FormatMySQLDateTime), ShouldEqual, "2019-11-21 04:10:40.24")
			So(t.UTC().Format(FormatMySQLDateTimeMilliseconds), ShouldEqual, "2019-11-21 04:10:40.240")
			So(t.UTC().Format(FormatMySQLDateTimeMicroseconds), ShouldEqual, "2019-11-21 04:10:40.240000")

			t = time.Unix(1574309440, 0)
			So(t.UTC().Format(FormatMySQLDateTime), ShouldEqual, "2019-11-21 04:10:40")
			So(t.UTC().Format(FormatMySQLDateTimeMilliseconds), ShouldEqual, "2019-11-21 04:10:40.000")
			So(t.UTC().Format(FormatMySQLDateTimeMicroseconds), ShouldEqual, "2019-11-21 04:10:40.000000")
		})

		Convey("Parse MySQl Date String To Time", func() {
			// mysql date string
			t, err := ParseDataMySQL("2017-07-04")
			So(err, ShouldBeNil)
			So(t.Year(), ShouldEqual, 2017)
			So(t.Month(), ShouldEqual, 7)
			So(t.Day(), ShouldEqual, 4)
			So(t.Hour(), ShouldEqual, 0)
			So(t.Minute(), ShouldEqual, 0)
			So(t.Second(), ShouldEqual, 0)
			name, offset := t.Zone()
			So(name, ShouldEqual, "UTC")
			So(offset, ShouldEqual, 0)
			So(t.UTC().Unix(), ShouldEqual, 1499126400)
			So(t.UTC().UnixNano(), ShouldEqual, 1499126400000000000)

			// invalid formats
			_, err = ParseDataMySQL("2017-07-04 03:32:31.987112")
			So(err, ShouldNotBeNil)
			_, err = ParseDataMySQL("12345")
			So(err, ShouldNotBeNil)
		})
	})
}
