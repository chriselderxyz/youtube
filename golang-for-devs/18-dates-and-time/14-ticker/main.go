package main

import (
	"fmt"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Ticker
/////////////////////////////////////////////////////////////

func main() {
	timer := time.NewTimer(3 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	// chan := time.Tick(time.Second)

	wg := sync.WaitGroup{}
	wg.Go(func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			case <-timer.C:
				fmt.Println("Timeout")
				return
			}
		}
	})
	wg.Wait()

	ticker.Stop()
	ticker.Reset(time.Second)
}
