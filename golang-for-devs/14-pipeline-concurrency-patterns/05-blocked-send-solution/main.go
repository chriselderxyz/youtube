package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Blocked Send Solution
/////////////////////////////////////////////////////////////

func main() {

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
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
				select {
				case <-ctx.Done():
					return
				case ch2 <- val * val:
					fmt.Println("Received")
				}
			}
		}
	})

	val := <-ch2
	fmt.Println("First Output: ", val)
	fmt.Println("Stop reading from channel 2")

	wg.Wait()
}
