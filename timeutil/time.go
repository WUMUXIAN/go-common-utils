package timeutil

import (
	"strings"
	"time"
)

const FormatMySQLDate = "2006-01-02"
const FormatMySQLDateTime = "2006-01-02 15:04:05.999999"
const FormatMySQLDateTimeMilliseconds = "2006-01-02 15:04:05.000"
const FormatMySQLDateTimeMicroseconds = "2006-01-02 15:04:05.000000"

// ParseDateTimeISO8601 parses a ISO8601 formatted datetime string into time.
func ParseDateTimeISO8601(dateTimeStr string) (t time.Time, err error) {
	// convert iso-8601 into rfc-3339 format
	rfc3339t := strings.Replace(dateTimeStr, " ", "T", 1) + "Z"

	// parse rfc-3339 datetime
	return time.Parse(time.RFC3339, rfc3339t)
}

// ParseDateTimeMySQL parses a MySQL formatted datetime string into time.
func ParseDateTimeMySQL(mySQLDateTimeStr string) (t time.Time, err error) {
	return time.Parse(FormatMySQLDateTime, mySQLDateTimeStr)
}

// ParseDateMySQL parses a MySQL formatted date string into time.
func ParseDateMySQL(mySQLDateStr string) (t time.Time, err error) {
	return time.Parse(FormatMySQLDate, mySQLDateStr)
}

// ParseTimeStamp parses a timestamp given in seconds, milliseconds or nanoseconds into time.
func ParseTimeStamp(timeStamp int64) (t time.Time) {
	if timeStamp < 9999999999 { // timestamp in seconds
		t = time.Unix(timeStamp, 0)
	} else if timeStamp < 9999999999999 { // timestamp in milli seconds
		t = time.Unix(0, timeStamp*int64(time.Millisecond))
	} else if timeStamp < 9999999999999999 { // timestamp in micro seconds
		t = time.Unix(0, timeStamp*int64(time.Microsecond))
	} else { // timestamp in nano seconds
		t = time.Unix(0, timeStamp)
	}
	return
}

// NowUTC returns current UTC time with optional duration addition
func NowUTC(addDuration ...time.Duration) time.Time {
	now := time.Now()

	if len(addDuration) == 1 {
		now = now.Add(addDuration[0])
	}

	return now.UTC()
}
