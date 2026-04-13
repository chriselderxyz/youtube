package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Channels - Closing Channels
/////////////////////////////////////////////////////////////

func Generate(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- rand.IntN(100) + 1
		time.Sleep(500 * time.Millisecond)
	}
}

func Transform(ch <-chan int) {
	for {
		val, ok := <-ch
		if ok == false {
			fmt.Println("Channel Closed")
			break
		}
		val = val * 10
		fmt.Println(val)
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Go(func() {
		Generate(ch)
	})

	wg.Go(func() {
		Transform(ch)
	})

	wg.Wait()
}
