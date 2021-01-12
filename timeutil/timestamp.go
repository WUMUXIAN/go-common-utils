// Package timeutil contains some utility functions regarding time operations
package timeutil

import (
	"math"
	"strconv"
	"sync/atomic"
	"time"
)

var (
	uniqueOrderedUint64Counter = int32(0)
)

// CurrentTimeStampStr gets current timestamp in string
func CurrentTimeStampStr(nano ...bool) string {
	return strconv.FormatInt(CurrentTimeStamp(nano...), 10)
}

// CurrentTimeStamp gets current timestamp in int64
func CurrentTimeStamp(nano ...bool) int64 {
	t := time.Now()
	timeStamp := t.Unix()
	if len(nano) > 0 && nano[0] {
		timeStamp = t.UnixNano()
	}
	return timeStamp
}

// GetTimeStamp get timestamp in int64 by year, month and day.
func GetTimeStamp(year, month, day int, nano ...bool) int64 {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	timeStamp := t.Unix()
	if len(nano) > 0 && nano[0] {
		timeStamp = t.UnixNano()
	}
	return timeStamp
}

// GetCurrentTimeStampMilli gets current timestamp in millis second in int64
func GetCurrentTimeStampMilli() int64 {
	return GetTimeStampFromTime(time.Now())
}

// UniqueIDAcrossProcessNow returns a unique ID across all processes.
func UniqueIDAcrossProcessNow() int64 {
	processID := math.MaxInt16
	timeStampSecond := time.Now().UnixNano() / 1e9 << 32
	extra := int16(atomic.AddInt32(&uniqueOrderedUint64Counter, 1))
	return timeStampSecond + int64(processID)<<16 + int64(extra)
}

func ParseToTimeStamp(dateTimeStr string) (timeStamp int64, err error) {
	t, err := ParseDateTimeISO8601(dateTimeStr)
	timeStamp = GetTimeStampFromTime(t)
	return
}

// GetStartOfNatureWeek get the start time of the nature week the given time stamp belongs to
func GetStartOfNatureWeek(timeStamp int64) (t time.Time) {
	date := time.Unix(timeStamp, 0)
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	for start.Weekday() != time.Monday { // iterate back to Monday
		start = start.AddDate(0, 0, -1)
	}
	return start
}

// GetEndOfNatureWeek get the end time of the nature week the given time stamp belongs to
func GetEndOfNatureWeek(timeStamp int64) (t time.Time) {
	date := time.Unix(timeStamp, 0)
	end := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, time.UTC)
	for end.Weekday() != time.Sunday { // iterate to Sundy
		end = end.AddDate(0, 0, 1)
	}
	return end
}

// GetStartOfNatureMonth get the start time of the nature month the given time stamp belongs to
func GetStartOfNatureMonth(timeStamp int64) (t time.Time) {
	date := time.Unix(timeStamp, 0)
	start := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	return start
}

// GetEndOfNatureMonth get the end time of the nature month the given time stamp belongs to
func GetEndOfNatureMonth(timeStamp int64) (t time.Time) {
	date := time.Unix(timeStamp, 0)
	end := time.Date(date.Year(), date.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	end = end.Add(-time.Nanosecond)
	return end
}

// Get timestamp from time.Time in millis second in int64
func GetTimeStampFromTime(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// AddMonths returns the time corresponding to adding the given number of months to t.
// For example, AddMonths(t, 1) applied to January 1, 2011 returns February 1, 2010.
//
// AddMonths does not normalizes its result as per AddDate. So, for example, adding one month to October 31 yields November 30.
// (AddDate returns Normalized form of November 31 which is December 1)
func AddMonths(t time.Time, d int) time.Time {
	newDate := t.AddDate(0, d, 0)
	if newDate.Day() == t.Day() { // No normalization happened
		return newDate
	}

	// When normalization happens one month is added always. We remove it first.
	newDate = newDate.AddDate(0, -1, 0)

	newYear := newDate.Year()
	newMonth := newDate.Month()
	newDay := newDate.Day()

	switch newMonth { // normalization happens when days overflow to next month. so we check last day of months
	case time.February:
		if t.Day() > 29 && newYear%4 == 0 {
			newDay = 29
		} else if t.Day() > 28 {
			newDay = 28
		}
		break
	case time.April, time.June, time.September, time.November:
		if t.Day() > 30 {
			newDay = 30
		}
	}

	return time.Date(newYear, newMonth, newDay, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}
