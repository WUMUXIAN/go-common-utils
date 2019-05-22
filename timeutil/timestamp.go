// Package timeutil contains some utility functions regarding time operations
package timeutil

import (
	"math"
	"strconv"
	"sync/atomic"
	"time"
)

const formatMySQLDateTime = "2006-01-02 15:04:05.999999"

var (
	uniqueOrderedUint64Counter = int32(0)
)

// CurrentTimeStampStr gets current timestamp in string
func CurrentTimeStampStr() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// CurrentTimeStamp gets current timestamp in int64
func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}

// GetTimeStamp get timestamp in int64 by year, month and day.
func GetTimeStamp(year, month, day int) int64 {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Unix()
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

func ParseToTimeStamp(dataTimeStr string) (timeStamp int64, err error) {
	t, err := ParseDataTimeISO8601(dataTimeStr)
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
