package random

import "time"

// @now
func Now() time.Time {
	return time.Now()
}

// @timestamp
func Timestamp() int64 {
	return time.Now().Unix()
}

// @datetime
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
