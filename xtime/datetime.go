package xtime

// Utils for datetime

import (
	"time"
)

// Get UTC Timestamp
func NowTS() int {
	t := time.Now()
	return int(t.Unix())
}

// Get Timestamp
func NowMsTS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Get UTC Timestamp
func GetTS(t time.Time) int {
	return int(t.Unix())
}

// Get the Date num string
func GetDateNumStr(t time.Time) string {
	return t.Format("20060102")
}

// Get the Date Hour num string
func GetDateHourStr(t time.Time) string {
	return t.Format("2006010215")
}

func TimeFromTS(ts int) time.Time {
	return time.Unix(int64(ts), 0)
}

// GetStrTime 获取当前时间  时间格式:2006-01-02 03:04:05
func GetStrTime() string {
	timeStamp := time.Now().Unix()
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 03:04:05")
}
