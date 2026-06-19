package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// Time Utility Functions
/////////////////////////////////////////////////////////////

func main() {
	time.Sleep(time.Second)

	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("AfterFunc")
	})

	timer.Reset(time.Second)

	ch := time.After(time.Second)
	current := <-ch
	fmt.Println("time.After channel signal: ", current)

	since := time.Since(time.Now().Add(-time.Hour))
	fmt.Println("time.Since(...): ", since)

	until := time.Until(time.Now().Add(time.Hour))
	fmt.Println("time.Until(...): ", until)
}
