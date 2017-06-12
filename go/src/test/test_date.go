package main

import (
	"fmt"
	"time"
	"utils/mdate"
)

func main() {
	var t int64 = DateUtils.GetNowUnixTime()
	fmt.Println("当前时间戳：", t)

	now := DateUtils.GetTimeByUnix(t) // time.Time类型，
	fmt.Println(now)                  // time.Time.String()方法：2017-05-24 20:41:14 +0800 CST

	month := now.Month() // Month类型，String方法返回 英文月份， 转成int方法： int(month)
	fmt.Println(now.Year(), int(month), now.Day())
	fmt.Println(now.Hour(), now.Minute(), now.Second())

	//month = time.Month(1)

	fmt.Println(now.Date()) // 年、月、日

	fmt.Println(DateUtils.DefaultDate(now))
	fmt.Println(DateUtils.DefaultTime(now))
	fmt.Println(DateUtils.Format(now, "yyyy-MM-dd HH"))

	fmt.Println(DateUtils.ParseDate("2017-05-24"))
	fmt.Println(DateUtils.ParseTime("2017-05-24 20:46:12"))
	fmt.Println(DateUtils.Parse("2017-05-24 20", "yyyy-MM-dd HH"))

	then := time.Date(2017, 05, 23, 20, 34, 58, 0, time.UTC)

	fmt.Println(then)

	fmt.Println(now.After(then))

	diff := now.Sub(then)

	fmt.Println(diff.Hours())

	m1, e := time.ParseDuration("-1h")
	fmt.Println(m1, e)
	y := now.Add(m1 * -3)
	fmt.Println(y)

	fmt.Println(DateUtils.AddDay(now, 3))
	fmt.Println(DateUtils.AddDay(now, -3))
}
