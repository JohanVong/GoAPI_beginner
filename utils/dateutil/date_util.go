package dateutil

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:04Z"
	apiDbLayout   = "2006-01-02 15:04:04"
)

// GetNow returns a current UTC time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns a current UTC time in string format
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat returns a current UTC time in DB format
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
