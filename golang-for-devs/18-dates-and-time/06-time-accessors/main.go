package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Time - Accessors
/////////////////////////////////////////////////////////////

func main() {
	t := time.Now()

	// Year, Month, Day
	y1 := t.Year()  // int
	m1 := t.Month() // time.Month
	dow := t.Weekday()
	dom := t.Day()
	doy := t.YearDay()
	y2, m2, d2 := t.Date()

	fmt.Println("t.Year(): ", y1)
	fmt.Println("t.Month(): ", m1)
	fmt.Println("t.Weekday(): ", dow)
	fmt.Println("t.Day(): ", dom)
	fmt.Println("t.YearDay(): ", doy)
	fmt.Println("t.Date() -> year: ", y2, " month: ", m2, " day: ", d2)

	// Hour, Minute, Second, Nanosecond
	h1 := t.Hour()
	min1 := t.Minute()
	sec1 := t.Second()
	h2, min2, sec2 := t.Clock()
	nano := t.Nanosecond()

	fmt.Println("t.Hour(): ", h1)
	fmt.Println("t.Minute(): ", min1)
	fmt.Println("t.Second(): ", sec1)
	fmt.Println("t.Clock() -> hour: ", h2, " min: ", min2, " second: ", sec2)
	fmt.Println("t.Nanosecond(): ", nano)

	// Time Zone
	loc := t.Location()
	name, offset := t.Zone()
	fmt.Println("t.Location(): ", loc)
	fmt.Println("t.Zone() -> name: ", name, " offset: ", offset)

	// Unix Timestamps
	unixSec := t.Unix()
	unixMilli := t.UnixMilli()
	unixMicro := t.UnixMicro()
	unixNano := t.UnixNano()

	fmt.Println("t.Unix(): ", unixSec)
	fmt.Println("t.UnixMilli(): ", unixMilli)
	fmt.Println("t.UnixMicro(): ", unixMicro)
	fmt.Println("t.UnixNano(): ", unixNano)
}
