package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

/////////////////////////////////////////////////////////////
// Cancel or Receive (OrDone)
/////////////////////////////////////////////////////////////

func CancelOrReceiveOfficial[T any](ctx context.Context, in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}

				if !CancelOrSend(ctx, out, val) {
					return
				}
			}
		}
	}()

	return out
}

func CancelOrReceiveBetter[T any](ctx context.Context, in <-chan T) (val T, ok bool) {
	select {
	case <-ctx.Done():
		return // 0, false
	case val, ok = <-in:
		return // val, true | 0, false
	}
}

func CancelOrSend[T any](ctx context.Context, ch chan<- T, val T) bool {
	select {
	case <-ctx.Done():
		return false
	case ch <- val:
		return true
	}
}

func Generate(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			if !CancelOrSend(ctx, out, rand.IntN(100)+1) {
				return
			}
		}
	}()

	return out
}

func Square(ctx context.Context, in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		for val := range CancelOrReceiveOfficial(ctx, in) {
			sq := val * val
			fmt.Println("Squaring Value: ", val, sq)
			time.Sleep(500 * time.Millisecond)

			if !CancelOrSend(ctx, out, sq) {
				return
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
			val, ok := CancelOrReceiveBetter(ctx, in)
			if !ok {
				return
			}

			dubs := val * 2
			fmt.Println("Doubling Value: ", val, dubs)
			time.Sleep(500 * time.Millisecond)

			if !CancelOrSend(ctx, out, dubs) {
				return
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
		fmt.Println("Workflow Output: ", val)
	}
}
