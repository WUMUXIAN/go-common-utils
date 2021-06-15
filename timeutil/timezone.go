package timeutil

import "time"

func GetTimezoneLocation(timezone string) (*time.Location, error) {
	if timezone == "" {
		timezone = "UTC"
	}

	return time.LoadLocation(timezone)
}
