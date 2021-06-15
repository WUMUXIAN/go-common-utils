package timeutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTimeUtilTimezone(t *testing.T) {
	Convey("Test Time Utilities For Timezone", t, func() {
		Convey("GetTimezoneLocation Should Work With Empty Timezone", func() {
			loc, err := GetTimezoneLocation("")
			So(err, ShouldBeNil)
			So(loc.String(), ShouldEqual, "UTC")
		})

		Convey("GetTimezoneLocation Should Work With UTC Timezone", func() {
			loc, err := GetTimezoneLocation("UTC")
			So(err, ShouldBeNil)
			So(loc.String(), ShouldEqual, "UTC")
		})

		Convey("GetTimezoneLocation Should Work With Europe/Zurich Timezone", func() {
			loc, err := GetTimezoneLocation("Europe/Zurich")
			So(err, ShouldBeNil)
			So(loc.String(), ShouldEqual, "Europe/Zurich")
		})

		Convey("GetTimezoneLocation Should Work With Asia/Singapore Timezone", func() {
			loc, err := GetTimezoneLocation("Asia/Singapore")
			So(err, ShouldBeNil)
			So(loc.String(), ShouldEqual, "Asia/Singapore")
		})

		Convey("GetTimezoneLocation Should Not Work With Invalid Timezone", func() {
			loc, err := GetTimezoneLocation("Asia/Hello")
			So(err, ShouldNotBeNil)
			So(loc.String(), ShouldEqual, "UTC")

			loc, err = GetTimezoneLocation("323232323")
			So(err, ShouldNotBeNil)
			So(loc.String(), ShouldEqual, "UTC")

			loc, err = GetTimezoneLocation("dsdjskdjsjd")
			So(err, ShouldNotBeNil)
			So(loc.String(), ShouldEqual, "UTC")
		})
	})
}
