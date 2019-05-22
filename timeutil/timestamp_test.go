package timeutil

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeUtilTimestamp(t *testing.T) {
	Convey("Test Time Utilities For Timestamp", t, func() {
		Convey("CurrentTimeStampStr Should Work", func() {
			So(CurrentTimeStampStr(), ShouldHaveLength, 10)
		})

		Convey("CurrentTimeStamp Should Work", func() {
			So(CurrentTimeStamp(), ShouldBeGreaterThan, 946684800)
		})

		Convey("GetTimeStamp Should Work", func() {
			So(GetTimeStamp(2010, 1, 1), ShouldEqual, 1262304000)
		})

		Convey("GetCurrentTimeStampMilli Should Work", func() {
			So(GetCurrentTimeStampMilli(), ShouldBeGreaterThan, 1262304000000)
		})

		Convey("UniqueIDAcrossProcessNow Should Work", func() {
			ids := make(chan int64, 3)
			for i := 0; i < 3; i++ {
				go func() {
					ids <- UniqueIDAcrossProcessNow()
				}()
			}
			id1 := <-ids
			id2 := <-ids
			id3 := <-ids
			So(id1, ShouldNotEqual, id2)
			So(id1, ShouldNotEqual, id3)
			So(id2, ShouldNotEqual, id3)
		})

		Convey("Parse ISO8601 DateTime String To Timestamp", func() {
			timestamp, err := ParseToTimeStamp("2017-07-04 03:32:31")
			So(err, ShouldBeNil)
			So(timestamp, ShouldEqual, 1499139151000)

			_, err = ParseToTimeStamp("12345")
			So(err, ShouldNotBeNil)
		})

		Convey("Get Start And End Of Nature Week Should Work", func() {
			t := GetStartOfNatureWeek(1538546431)
			So(t.String(), ShouldEqual, "2018-10-01 00:00:00 +0000 UTC")
			t = GetEndOfNatureWeek(1538546431)
			So(t.String(), ShouldEqual, "2018-10-07 23:59:59.999999999 +0000 UTC")

			t = GetStartOfNatureWeek(1489816819)
			So(t.String(), ShouldEqual, "2017-03-13 00:00:00 +0000 UTC")
			t = GetEndOfNatureWeek(1489816819)
			So(t.String(), ShouldEqual, "2017-03-19 23:59:59.999999999 +0000 UTC")
		})

		Convey("Get Start And End Of Nature Month Should Work", func() {
			t := GetStartOfNatureMonth(1538546431)
			So(t.String(), ShouldEqual, "2018-10-01 00:00:00 +0000 UTC")
			t = GetEndOfNatureMonth(1538546431)
			So(t.String(), ShouldEqual, "2018-10-31 23:59:59.999999999 +0000 UTC")

			t = GetStartOfNatureMonth(1489816819)
			So(t.String(), ShouldEqual, "2017-03-01 00:00:00 +0000 UTC")
			t = GetEndOfNatureMonth(1489816819)
			So(t.String(), ShouldEqual, "2017-03-31 23:59:59.999999999 +0000 UTC")
		})
	})
}
