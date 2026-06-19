package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Time - Operations
/////////////////////////////////////////////////////////////

func main() {
	t1 := time.Now()

	// Add and subract
	t2 := t1.Add(time.Hour * 2)
	t3 := t1.AddDate(10, 11, 12)

	t4 := t2.Sub(t3) // t2 - t3

	// Rounding
	t5 := t1.Round(time.Hour)
	t6 := t1.Truncate(time.Hour)

	// Timezones
	t7 := t1.UTC()
	t8 := t1.Local()
	loc, _ := time.LoadLocation("Europe/London")
	t9 := t1.In(loc)

	// Functions that Strip Monotonic Clock:
	// AddDate
	// Round
	// Truncate
	// UTC
	// Local
	// In

	fmt.Println("t1 - time.Now(): ", t1)
	fmt.Println("t2 - t1.Add(time.Hour*2): ", t2)
	fmt.Println("t3 - t1.AddDate(10, 11, 12): ", t3)
	fmt.Println("t4 - t2.Sub(t3): ", t4)
	fmt.Println("t5 - t1.Round(time.Hour): ", t5)
	fmt.Println("t6 - t1.Truncate(time.Hour): ", t6)
	fmt.Println("t7 - t1.UTC(): ", t7)
	fmt.Println("t8 - t1.Local(): ", t8)
	fmt.Println("t9 - t1.In('Europe/London'): ", t9)

}
