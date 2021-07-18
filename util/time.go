package util

import (
	"strconv"
	"time"
)

// GetNowTime get now time(2006-01-02 15:04:05)
func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetNowTime2 get now time(2006-01-02 15:04:05)
func GetNowTime2() string {
	return time.Now().Format("20060102")
}
func GetNowTime3() string {
	return time.Now().Format("20060102150405")
}

// GetCurrentTime get current time
func GetCurrentTime() time.Time {
	return time.Now()
}

// GetNowTimeUnix get now time (毫秒级)
func GetNowTimeUnix() int64 {
	return time.Now().UnixNano() / 1e6
}

// TimeUnixToString
func TimeUnixToString(t int64) string {
	t = t / 1e3
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

// GetNowTimeUnixNanoString
func GetNowTimeUnixNanoString() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// FormatTimeStr  format time to str
func FormatTimeStr(timeStr string) (string, error) {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", timeStr, loc)
	return theTime.Format("2006/01/02 15:04:05"), err
}
