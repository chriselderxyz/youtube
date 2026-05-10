package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Blocked Send Problem
/////////////////////////////////////////////////////////////

func main() {
	
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Go(func() {
		defer close(ch1)
		for {
			select {
			case <-ctx.Done():
				return
			case ch1 <- rand.IntN(100) + 1:
			}
		}
	})

	wg.Go(func() {
		defer close(ch2)
		for {
			select {
			case <-ctx.Done():
				return
			case val := <-ch1:
				fmt.Println("Sending to ch2")
				ch2 <- val * val // Will be blocked and cause deadlock
				fmt.Println("Received")
			}
		}
	})

	val := <-ch2
	fmt.Println("First Output: ", val)
	fmt.Println("Stop reading from channel 2")

	wg.Wait()
}
