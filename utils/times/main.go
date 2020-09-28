package times

import (
	"time"
)

const (
	DefaultTimeFormat    string = "2006-01-02 15:04:05"
	DateFormatEN         string = "2006/01/02 15:04:05"
	DateFormatNoSpan     string = "20060102150405"
	DateFormatLongNoSpan string = "20060102150405000"
)

var CstZone = time.FixedZone("CST", 8*3600)

//毫秒数 转 Time
func TimestampToTime(timestamp int64) time.Time {
	ns := timestamp % 1000 * 1e6
	return time.Unix(timestamp/1000, ns)
}

//毫秒数 转 字符串
func TimestampToStr(timestamp int64, format string) string {
	return TimeToStr(TimestampToTime(timestamp), format)
}

//Time 转 毫秒数
func TimeToTimestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

//Time 转 字符串
func TimeToStr(t time.Time, format string) string {
	if format == "" {
		format = DefaultTimeFormat
	}
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), CstZone).Format(format)
}

//字符串 转 Time
func StrToTime(timeStr string, format string) (t time.Time, err error) {
	if format == "" {
		format = DefaultTimeFormat
	}
	return time.ParseInLocation(format, timeStr, CstZone) //使用模板在对应时区转化为time.time类型
}

//字符串 转 Time Must
func StrToTimeMust(timeStr string, format string) time.Time {
	t, _ := StrToTime(timeStr, format)
	return t
}

//字符串 转 Timestamp
func StrToTimestamp(timeStr string, format string) (ts int64, err error) {
	t, err := StrToTime(timeStr, format)
	if err != nil {
		return 0, err
	}
	return TimeToTimestamp(t), nil
}

//字符串 转 Timestamp Must
func StrToTimestampMust(timeStr string, format string) (ts int64) {
	return TimeToTimestamp(StrToTimeMust(timeStr, format))
}
