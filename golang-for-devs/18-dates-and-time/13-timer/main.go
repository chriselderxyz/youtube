package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Timer
/////////////////////////////////////////////////////////////

func main() {
	timer := time.NewTimer(time.Second * 3)

	ch := make(chan int)
	select {
	case <-ch:
	case <-timer.C:
		fmt.Println("Timeout")
	}

	fired := timer.Stop()
	fmt.Println("Fired: ", fired)

	active := timer.Reset(time.Second * 1)
	fmt.Println("Active: ", active)
}
