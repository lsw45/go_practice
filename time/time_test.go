package time

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {
	var t11 time.Time
	fmt.Println(t11)          //0001-01-01 00:00:00 +0000 UTC
	fmt.Println(t11.IsZero()) //true
	timeFormat()
	// timeLocation()
	// timeUnix()
	// timeDuration()
	// timeDiff()
	// timeTicker()
	timePointerNull()
}

func timeFormat() {
	fmt.Println("--------------------------------timeFormat-----------------------------------")

	var t = new(time.Time)
	*t, _ = time.Parse("2006-01-02 15:04:05", "2018-04-23 12:24:51")
	fmt.Println("Parse:", *t)

	*t = time.Now()
	fmt.Printf("time.Now():%T,%v\n", *t, *t)
	fmt.Println("time.Now().String():", t.String())
	fmt.Println("time.Now().Format:", t.Format("2006-01-02 15:04:05"))
	/*
		Parse: 2018-04-23 12:24:51 +0000 UTC
		time.Now():time.Time,2019-10-11 23:21:15.1305433 +0800 CST m=+0.007977801
		time.Now().String(): 2019-10-11 23:21:15.1305433 +0800 CST m=+0.007977801
		time.Now().Format: 2019-10-11 23:21:15
	*/
	fmt.Println("--------------------------------end-----------------------------------")
}

func timeLocation() {
	fmt.Println("--------------------------------timeLocation-----------------------------------")

	var t = new(time.Time)
	*t = time.Now()
	*t, _ = time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:04:05"), time.Local)
	fmt.Println("time.ParseInLocation():", *t)

	loc, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println("time.Now().In(loc):", t.In(loc))

	fmt.Printf("相同时间不同时区的时间戳差值:%d\n", t.In(loc).Unix()-t.Unix())

	// 默认UTC
	loc, _ = time.LoadLocation("")
	fmt.Println(loc)

	// 服务器设定的时区，一般为CST
	loc, _ = time.LoadLocation("Local")
	fmt.Println(loc)

	// 美国洛杉矶PDT
	loc, _ = time.LoadLocation("America/Los_Angeles")
	fmt.Println(loc)

	fmt.Println("--------------------------------end-----------------------------------")
}

func timeUnix() {
	fmt.Println("--------------------------------timeUnix-----------------------------------")

	var t = new(time.Time)
	*t = time.Now()

	fmt.Println(t.Unix())

	dt, _ := time.Parse("2006-01-02 15:04:05", t.Format("2006-01-02 15:04:05"))
	fmt.Printf("%T,%v\n", dt, dt)
	fmt.Println(dt.Unix())

	fmt.Println("--------------------------------end-----------------------------------")
}

func timeDuration() {
	fmt.Println("--------------------------------timeDuration-----------------------------------")

	tp, _ := time.ParseDuration("4h28m32.88888888885s")
	fmt.Println(tp.Truncate(1000), tp.Seconds(), tp.Nanoseconds()) //4h28m32.888888s 16112.888888888 16112888888888
	fmt.Printf("I've got %.1f hours of work left.\n", tp.Hours())
	fmt.Printf("The movie is %.0f minutes long.\n", tp.Minutes())
	fmt.Printf("one microsecond has %d nanoseconds.\n", tp.Nanoseconds())
	fmt.Printf("take off in t-%.0f seconds.\n", tp.Seconds())
	/*
		I've got 4.5 hours of work left.
		The movie is 269 minutes long.
		one microsecond has 16112888888888 nanoseconds.
		take off in t-16113 seconds.
	*/

	dur, _ := time.ParseDuration("1h30m35.88888888888888888s")
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, v := range round {
		fmt.Printf("Round:%6s:%s\n", v, dur.Round(v).String())
	}
	for _, v := range round {
		fmt.Printf("Truncate:%6s:%s\n", v, dur.Truncate(v).String())
	}
	/*
		Round:   1ns:1h30m35.888888888s
		Round:   1µs:1h30m35.888889s
		Round:   1ms:1h30m35.889s
		Round:    1s:1h30m36s
		Round:    2s:1h30m36s
		Round:  1m0s:1h31m0s
		Round: 10m0s:1h30m0s
		Round:1h0m0s:2h0m0s
		Truncate:   1ns:1h30m35.888888888s
		Truncate:   1µs:1h30m35.888888s
		Truncate:   1ms:1h30m35.888s
		Truncate:    1s:1h30m35s
		Truncate:    2s:1h30m34s
		Truncate:  1m0s:1h30m0s
		Truncate: 10m0s:1h30m0s
		Truncate:1h0m0s:1h0m0s
	*/
	fmt.Println("--------------------------------end-----------------------------------")
}

func timeDiff() {
	fmt.Println("--------------------------------timeDiff-----------------------------------")

	now := time.Now()
	t, _ := time.Parse("2006-01-02 15:04:05", "2018-04-23 12:24:51")
	fmt.Println("Time.Sub(Time):", now.Sub(t))
	fmt.Println("Time.Add():", now.Add(time.Duration(10)*time.Minute))
	fmt.Println("Time.AddDate():", now.AddDate(1, 1, 1))
	/*
		Time.Sub(Time): 12867h45m23.158173s
		Time.Add(): 2019-10-12 00:20:14.158173 +0800 CST m=+600.004987101
		Time.AddDate(): 2020-11-13 00:10:14.158173 +0800 CST
	*/
	fmt.Println("Time.After(Time):", t.After(now))
	fmt.Println("Time.Before(Time):", t.Before(now))
	fmt.Println("Time.Equal(Time):", t.Equal(now))
	/*
		Time.After(Time): false
		Time.Before(Time): true
		Time.Equal(Time): false
	*/

	// 计算两个时间点的相差天数
	dt1 := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	dt2 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println(int(math.Ceil(dt1.Sub(dt2).Hours() / 24)))

	time.Sleep(time.Second * 3)
	fmt.Println("time.Since(Time):", time.Since(now))

	fmt.Println("--------------------------------end-----------------------------------")
}

func timeTicker() {
	fmt.Println("--------------------------------timeTicker-----------------------------------")

	// 可通过调用ticker.Stop取消
	ticker := time.NewTicker(1 * time.Second)
	fmt.Printf("%T\n", *ticker) //time.Ticker
	for i := 0; i < 10; i++ {
		t := <-ticker.C
		fmt.Println(t.String())
	}

	// 无法取消
	tick := time.Tick(1 * time.Second)
	fmt.Printf("%T\n", tick) //<-chan time.Time
	for v := range tick {
		fmt.Println(v)
	}
	fmt.Println("--------------------------------end-----------------------------------")
}

func timePointerNull() {
	p := struct {
		Tp *time.Time `json:"tp"`
	}{}
	fmt.Printf("%+v\n", p) // {Tp:<nil>}

	j, _ := json.Marshal(p)

	fmt.Println(string(j)) // {"tp":null}

	tt1 := time.Now()
	p.Tp = &tt1
	fmt.Printf("%+v\n", p) // {Tp:2019-10-25 16:24:23.4423372 +0800 CST m=+0.034962101}

	j, _ = json.Marshal(p)

	fmt.Println(string(j)) // {"tp":"2019-10-25T16:24:23.4423372+08:00"}

	p1 := struct {
		Tp *time.Time `json:"tp,omitempty"`
	}{}

	j, _ = json.Marshal(p1)

	fmt.Printf("omitempty:%s\n", string(j)) // omitempty:{}
}
