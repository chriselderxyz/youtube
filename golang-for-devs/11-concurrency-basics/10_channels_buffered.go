package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Channels - Buffered Channels
/////////////////////////////////////////////////////////////

func Generate(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("Queueing Data - Burst")
			ch <- rand.IntN(100) + 1
		}
		time.Sleep(3 * time.Second)
	}
}

func Transform(ch <-chan int) {
	for val := range ch {
		fmt.Println("Processing Data")
		time.Sleep(time.Duration(rand.IntN(2000)+1) * time.Millisecond)
		val = val * 10
		fmt.Println(val)
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 25)

	wg.Go(func() {
		Generate(ch)
	})

	wg.Go(func() {
		Transform(ch)
	})

	wg.Wait()
}
