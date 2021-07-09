package time

import (
	"time"
)

const (
	FormatStandardDate     = "2006-01-02"
	FormatStandardTime     = "15:04:05"
	FormatStandardDateTime = "2006-01-02 15:04:05"
	FormatDate8            = "20060102"
	FormatTime6            = "150405"
	FormatDateTime86       = FormatDate8 + FormatTime6
)

func Timestamp() int64 {
	return time.Now().Unix()
}

func Date() string {
	return time.Now().Format(FormatStandardDate)
}

func Time() string {
	return time.Now().Format(FormatStandardTime)
}

func DateTime() string {
	return time.Now().Format(FormatStandardDateTime)
}

func Date8() string {
	return time.Now().Format(FormatDate8)
}

func Time6() string {
	return time.Now().Format(FormatTime6)
}

func DateTime86() string {
	return time.Now().Format(FormatDateTime86)
}

func DayUnixFirst(unix int64) int64 {
	now := time.Unix(unix, 0)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
}

func DayUnixLast(unix int64) int64 {
	now := time.Unix(unix, 0)
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Unix()
}

func TodayUnixFirst() int64 {
	return DayUnixFirst(Timestamp())
}

func TodayUnixLast() int64 {
	return DayUnixLast(Timestamp())
}

func YesterdayUnixFirst() int64 {
	return DayUnixFirst(Timestamp() - 86400)
}

func YesterdayUnixLast() int64 {
	return DayUnixLast(Timestamp() - 86400)
}
