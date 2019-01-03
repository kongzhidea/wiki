package DateUtils

import (
	"fmt"
	"strings"
	"time"
)

// time.Now() 当前时间
// time.Time.String()方法：2017-05-24 20:41:14 +0800 CST
// time.Year(),Month(),Day(),Hour(),Minute(),Second() 返回年月日，时分秒；  Date() 返回（年、月、日）
// 指定日期：time.Date(2009, 2, 17, 20, 34, 58, 0, time.UTC)  // month可以传Month类型，也可以传int类型。

// 时间比较  Before(),After()， 判断相同可以根据时间戳来判断

// 时间比较： Sub() 返回Duration类型，可以获取相差小时，分钟，秒， 如果想获取相差天数，则根据秒除下就行
// 时间比较： Add() 参数为Duration类型，可以通过time.ParseDuration()来获取，如"1h","-1h"，也可以 3 * time.Second 获取

// 当前时间的unix时间戳，精确到秒
func GetNowUnixTime() int64 {
	return time.Now().Unix()
}

// 时间戳转Time类型
func GetTimeByUnix(t int64) time.Time {
	return time.Unix(t, 0)
}

func DefaultTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func DefaultDate(time time.Time) string {
	return time.Format("2006-01-02")
}

// yyyy-MM-dd HH:mm:ss
func Format(time time.Time, format string) string {
	format = strings.Replace(format, "yyyy", "2006", -1)
	format = strings.Replace(format, "MM", "01", -1)
	format = strings.Replace(format, "dd", "02", -1)
	format = strings.Replace(format, "HH", "15", -1)
	format = strings.Replace(format, "mm", "04", -1)
	format = strings.Replace(format, "ss", "05", -1)
	return time.Format(format)
}

func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", s, time.Local)
}

func ParseTime(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
}

func Parse(s, format string) (time.Time, error) {
	format = strings.Replace(format, "yyyy", "2006", -1)
	format = strings.Replace(format, "MM", "01", -1)
	format = strings.Replace(format, "dd", "02", -1)
	format = strings.Replace(format, "HH", "15", -1)
	format = strings.Replace(format, "mm", "04", -1)
	format = strings.Replace(format, "ss", "05", -1)
	return time.ParseInLocation(format, s, time.Local)
}

func AddDay(date time.Time, day int64) time.Time {
	delta, _ := time.ParseDuration(fmt.Sprintf("%dh", day*24))
	return date.Add(delta)
}
