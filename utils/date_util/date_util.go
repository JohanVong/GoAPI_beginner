package date_util

import "time"

const apiDateLayout = "2006-01-02T15:04Z"

// GetNow returns a current UTC time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns a current UTC time in string format
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
