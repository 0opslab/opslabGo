package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Hour())
	p(now.Minute())
	p(now.Second())
	p(now.Nanosecond())
	p(now.Location())
	p(now.Weekday())

	then := time.Now()
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	//方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//你可以使用 Add 将时间后移一个时间间隔，或者使用一个 - 来将时间前移一个时间间隔。
	p(then.Add(diff))
	p(then.Add(-diff))

	currentTime := time.Now().Local()
	newFormat := currentTime.Format("2006-01-02 15:04:05.000")
	fmt.Println(newFormat)

	newFormat1 := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println(newFormat1)

	//获取时间戳
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))

	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	fmt.Println(tm2.Unix())

	tm3, _ := time.Parse("2006-01-02 15:04:05", "2017-08-02 13:14:15")
	fmt.Println(tm3)

	tm4, _ := time.Parse("2006/01/02 15:04:05", "2017/08/02 13:14:15")
	fmt.Println(tm4)
}
