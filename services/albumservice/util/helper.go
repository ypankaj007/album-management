package util

import (
	"time"
)

// UnixMilli use to get milliseconds of given time
// @params - time
// return - milliseconds of the given time
func UnixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// CurrentTimeInMilli use to get current time in milliseconds
// This function will use when we need current timestamp
// This functions return current timestamp (in millisecods)
func CurrentTimeInMilli() int64 {
	return UnixMilli(time.Now())
}
