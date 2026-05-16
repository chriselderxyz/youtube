package main

import (
	"fmt"
	"sync"
)

/////////////////////////////////////////////////////////////
// sync.Cond
/////////////////////////////////////////////////////////////

func main() {
	wg := sync.WaitGroup{}

	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)

	counter := 0

	wg.Go(func() {
		for i := 0; i < 150; i++ {
			cond.L.Lock()
			counter++
			cond.Broadcast()
			cond.L.Unlock()
		}
	})

	wg.Go(func() {
		cond.L.Lock()

		for counter < 100 {
			cond.Wait()
		}

		cond.L.Unlock()
		fmt.Println("Counter >= 100", counter)
	})

	wg.Wait()
}
