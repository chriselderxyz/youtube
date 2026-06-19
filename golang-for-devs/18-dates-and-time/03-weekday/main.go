package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Weekday
/////////////////////////////////////////////////////////////

func main() {
	wd1 := time.Monday
	wd2 := time.Weekday(0) // Sunday

	fmt.Println("wd1 - time.Monday: ", wd1.String())
	fmt.Println("wd2 - time.Weekday(0): ", wd2.String())
}
