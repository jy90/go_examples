package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1. get current time
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)

	// 2. get special time
	t2 := time.Date(2019, 2, 15, 16, 30, 28, 0, time.Local)
	fmt.Println(t2)

	// 3. time --> string
	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)

	s2 := t1.Format("2006/01/02")
	fmt.Println(s2)

	// 4. string --> time
	s3 := "1999年10月01日"
	t3, err := time.Parse("2006年01月02日", s3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3)
	fmt.Printf("%T\n", t3)

	// 5. 根据当前时间获取 年月日时分秒
	fmt.Println(t1.String())
	year, month, day := t1.Date()
	fmt.Println(year, month, day)

	hour, minute, second := t1.Clock()
	fmt.Println(hour, minute, second)

	fmt.Println(t1.Year())
	fmt.Println(t1.YearDay())
	fmt.Println(t1.Month())
	fmt.Println(t1.Day())
	fmt.Println(t1.Hour())
	fmt.Println(t1.Minute())
	fmt.Println(t1.Second())
	fmt.Println(t1.Nanosecond())

	fmt.Println(t1.Weekday())

	// 6. 时间戳 指定时间的时间戳/当前时间的时间戳(秒)
	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	ts1 := t4.Unix()
	fmt.Println(ts1)

	ts2 := t1.Unix()
	fmt.Println(ts2)

	// 时间戳(纳秒)
	ts3 := t4.UnixNano()
	fmt.Println(ts3) // 3600 000 000 000

	ts4 := t1.UnixNano()
	fmt.Println(ts4)

	// 7. 时间间隔
	t5 := t1.Add(time.Minute)
	fmt.Println(t1)
	fmt.Println(t5)
	fmt.Println(t1.Add(24 * time.Hour))

	t6 := t1.AddDate(1, 0, 0)
	fmt.Println(t6)

	d1 := t5.Sub(t1)
	fmt.Println(d1)

	// 7. time.Sleep
	time.Sleep(3 * time.Second)
	fmt.Println("3s done.")

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1
	fmt.Println(randNum)
	time.Sleep(time.Duration(randNum) * time.Second)
	fmt.Println("rand time", randNum)
}
