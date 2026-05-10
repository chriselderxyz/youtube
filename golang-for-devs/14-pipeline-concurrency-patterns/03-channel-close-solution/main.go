package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

/////////////////////////////////////////////////////////////
// Closing Channel Solution
/////////////////////////////////////////////////////////////

func Generate(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case out <- rand.IntN(100) + 1:
			}
		}
	}()

	return out
}

func Square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:

				// Check if channel is closed
				if !ok {
					return
				}
				
				sq := val * val
				fmt.Println("Squaring Value: ", val, sq)
				time.Sleep(500 * time.Millisecond)
				out <- sq
			}
		}
	}()

	return out
}

func Double(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:

				// Check if channel is closed
				if !ok {
					return
				}

				dubs := val * 2
				fmt.Println("Doubling Value: ", val, dubs)
				time.Sleep(500 * time.Millisecond)
				out <- dubs
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch1 := Generate(ctx)
	ch2 := Square(ctx, ch1)
	ch3 := Double(ctx, ch2)

	for val := range ch3 {
		fmt.Println("Output: ", val)
	}
}
