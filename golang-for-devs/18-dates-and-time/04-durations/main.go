package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Duration
/////////////////////////////////////////////////////////////

// const (
// 	Nanosecond  Duration = 1
// 	Microsecond          = 1000 * Nanosecond
// 	Millisecond          = 1000 * Microsecond
// 	Second               = 1000 * Millisecond
// 	Minute               = 60 * Second
// 	Hour                 = 60 * Minute
// )

func main() {
	d := time.Hour * 24

	// Query
	hours := d.Hours()
	fmt.Println("Hours: ", hours)

	mins := d.Minutes()
	secs := d.Seconds()
	fmt.Println("Minutes: ", mins)
	fmt.Println("Seconds: ", secs)

	milli := d.Milliseconds()
	micro := d.Microseconds()
	nano := d.Nanoseconds()
	fmt.Println("Milli: ", milli)
	fmt.Println("Micro: ", micro)
	fmt.Println("Nano: ", nano)

	// Operations
	abs := d.Abs()
	fmt.Println("Absolute: ", abs)

	rounded := d.Round(time.Second * 4)
	fmt.Println("Rounded: ", rounded)

	truncate := d.Truncate(time.Hour * 5) // -22 | -20
	fmt.Println("Truncated: ", truncate)

	// Stringer
	// XhoursXminXsec ...
	str := d.String()
	fmt.Println("Stringified: ", str)
}
