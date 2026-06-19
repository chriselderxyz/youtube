package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Time - Creating
/////////////////////////////////////////////////////////////

func main() {

	// Current date + time
	// Includes monotonic clock
	t1 := time.Now()

	// Custom date + time
	t2 := time.Date(
		2026,
		time.May,
		15,
		13,
		59,
		45,
		1000,
		time.Local,
	)

	// Parse from a string
	str3 := "Apr 10, 2026 at 10:26:33am (EST)"
	// first month, 2nd day, 3rd hour, 4th min, 5th second, 6th year, -7 timezone offset
	layoutStr := "Jan 2, 2006 at 3:04:05pm (MST)"
	t3, err := time.Parse(layoutStr, str3)

	// Parse using a built in format string
	rfcStr := "2042-11-07T03:26:51-05:00"
	t4, err := time.Parse(time.RFC3339, rfcStr)

	// Parse with custom time zone
	noTimeZone := "Apr 10, 2026, at 10:26am"
	layoutStr2 := "Jan 2, 2006, at 3:04pm"
	loc, _ := time.LoadLocation("Europe/London")
	t5, err := time.ParseInLocation(layoutStr2, noTimeZone, loc)

	// Time from Unix timestamp
	t6 := time.UnixMicro(int64(1779127200000000))
	t7 := time.UnixMilli(int64(1779127200000))
	t8 := time.Unix(int64(1779127200), int64(1000))

	fmt.Println("t1 - time.Now: ", t1)
	fmt.Println("t2 - time.Date: ", t2)
	fmt.Println("t3 - time.Parse: ", t3, err)
	fmt.Println("t4 - time.Parse: ", t4, err)
	fmt.Println("t5 - time.ParseInLocation: ", t5, err)
	fmt.Println("Unix Times: ", t6, t7, t8)
}
