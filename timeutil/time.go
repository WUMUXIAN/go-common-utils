package timeutil

import (
	"strings"
	"time"
)

const formatMySQLDateTime = "2006-01-02 15:04:05.999999"

// ParseDataTimeISO8601 parses a ISO8601 formated data time string into time.
func ParseDataTimeISO8601(dataTimeStr string) (t time.Time, err error) {
	// convert iso-8601 into rfc-3339 format
	rfc3339t := strings.Replace(dataTimeStr, " ", "T", 1) + "Z"

	// parse rfc-3339 datetime
	return time.Parse(time.RFC3339, rfc3339t)
}

// ParseDataTimeMySQL parses a MySQL formated data time string into time.
func ParseDataTimeMySQL(mySQLDateTimeStr string) (t time.Time, err error) {
	return time.Parse(formatMySQLDateTime, mySQLDateTimeStr)
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
