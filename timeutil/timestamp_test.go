package timeutil

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeUtil(t *testing.T) {
	Convey("Test Time Utilities", t, func() {
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
	})
}
