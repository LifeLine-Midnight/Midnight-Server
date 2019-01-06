package utils

import "time"

type TimeUtils struct{}

// TimestampToStr 将时间戳转化为 "2019-01-01 00:00:00" 的字符串
func (*TimeUtils) TimestampToStr(ts int64) string {
	var timeLayout = "2006-01-02 15:04:05"
	timeStr := time.Unix(ts, 0).Format(timeLayout)

	return timeStr
}

// StrToTimestamp 将 "2019-01-01 00:00:00" 格式的字符串转化为时间戳
func (*TimeUtils) StrToTimestamp(timeStr string) int64 {
	var timeLayout = "2006-01-02 15:04:05"[0:len(timeStr)]
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(timeLayout, timeStr, loc)

	return t.Unix()
}
