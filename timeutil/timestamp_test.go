package timeutil

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeUtilTimestamp(t *testing.T) {
	Convey("Test Time Utilities For Timestamp", t, func() {
		Convey("CurrentTimeStampStr Should Work", func() {
			So(CurrentTimeStampStr(), ShouldHaveLength, 10)
			So(CurrentTimeStampStr(true), ShouldHaveLength, 19)
		})

		Convey("CurrentTimeStamp Should Work", func() {
			So(CurrentTimeStamp(), ShouldBeGreaterThan, 946684800)
			So(CurrentTimeStamp(true), ShouldBeGreaterThan, 946684800000000000)
		})

		Convey("GetTimeStamp Should Work", func() {
			So(GetTimeStamp(2010, 1, 1), ShouldEqual, 1262304000)
			So(GetTimeStamp(2010, 1, 1, true), ShouldEqual, 1262304000000000000)
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

		Convey("Add Months To Time Should Work", func() {
			t := time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC)
			Convey("Add 1 Month To 1st Of March Should Return 1st Of April", func() {
				t2 := AddMonths(t, 1)
				So(t2.String(), ShouldEqual, "2021-04-01 00:00:00 +0000 UTC")
			})

			Convey("Subtract 1 Month From 1st of March Should Return 1st Of February", func() {
				t2 := AddMonths(t, -1)
				So(t2.String(), ShouldEqual, "2021-02-01 00:00:00 +0000 UTC")
			})

			Convey("Subtract 5 Months From 1st of March Should Return 1st Of October Last Year", func() {
				t2 := AddMonths(t, -5)
				So(t2.String(), ShouldEqual, "2020-10-01 00:00:00 +0000 UTC")
			})

			Convey("Add 5 Months To 1st of March Should Return 1st Of August", func() {
				t2 := AddMonths(t, 5)
				So(t2.String(), ShouldEqual, "2021-08-01 00:00:00 +0000 UTC")
			})

			Convey("Add 11 Months To 1st of March Should Return 1st Of February Next year", func() {
				t2 := AddMonths(t, 11)
				So(t2.String(), ShouldEqual, "2022-02-01 00:00:00 +0000 UTC")
			})

			Convey("Add 15 Months To 1st of March Should Return 1st Of June Next year", func() {
				t2 := AddMonths(t, 15)
				So(t2.String(), ShouldEqual, "2022-06-01 00:00:00 +0000 UTC")
			})

			t = time.Date(2021, 3, 31, 0, 0, 0, 0, time.UTC)
			Convey("Add 1 Month To 31st Of March Should Return 30th Of April", func() {
				t2 := AddMonths(t, 1)
				So(t2.String(), ShouldEqual, "2021-04-30 00:00:00 +0000 UTC")
			})

			Convey("Subtract 1 Month From 31st of March Should Return 28th Of February", func() {
				t2 := AddMonths(t, -1)
				So(t2.String(), ShouldEqual, "2021-02-28 00:00:00 +0000 UTC")
			})

			Convey("Subtract 13 Months From 31st of March Should Return 29th Of February", func() {
				t2 := AddMonths(t, -13)
				So(t2.String(), ShouldEqual, "2020-02-29 00:00:00 +0000 UTC")
			})

			Convey("Subtract 5 Months From 31st of March Should Return 31st Of October Last Year", func() {
				t2 := AddMonths(t, -5)
				So(t2.String(), ShouldEqual, "2020-10-31 00:00:00 +0000 UTC")
			})

			Convey("Add 5 Months To 31st of March Should Return 31st Of August", func() {
				t2 := AddMonths(t, 5)
				So(t2.String(), ShouldEqual, "2021-08-31 00:00:00 +0000 UTC")
			})

			Convey("Add 11 Months To 31st of March Should Return 28th Of February Next year", func() {
				t2 := AddMonths(t, 11)
				So(t2.String(), ShouldEqual, "2022-02-28 00:00:00 +0000 UTC")
			})

			Convey("Add 15 Months To 31st of March Should Return 30th Of June Next year", func() {
				t2 := AddMonths(t, 15)
				So(t2.String(), ShouldEqual, "2022-06-30 00:00:00 +0000 UTC")
			})

			Convey("Add 35 Months To 31st of March Should Return 29th Of February Next Leap year", func() {
				t2 := AddMonths(t, 35)
				So(t2.String(), ShouldEqual, "2024-02-29 00:00:00 +0000 UTC")
			})

			t = time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC)
			Convey("Add 12 Months To 29th of February Should Return 28th Of February Next year", func() {
				t2 := AddMonths(t, 12)
				So(t2.String(), ShouldEqual, "2021-02-28 00:00:00 +0000 UTC")
			})
		})
	})
}
