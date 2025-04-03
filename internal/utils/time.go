package utils

import "time"

// TimestampToTime 时间戳转时间
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// TimeToStr 时间转字符串
func TimeToStr(nowTime time.Time) string {
	return nowTime.Format("2006-01-02 15:04:05")
}

// StrToTime 字符串转时间
func StrToTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

// StrToTimestamp 字符串转时间戳
func StrToTimestamp(str string) (int64, error) {
	time, err := StrToTime(str)
	return time.Unix(), err
}

// TimestampToStr 时间戳转字符串
func TimestampToStr(timestamp int64) string {
	return TimestampToTime(timestamp).Format("2006-01-02 15:04:05")
}
