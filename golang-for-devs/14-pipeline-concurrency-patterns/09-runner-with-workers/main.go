package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// ///////////////////////////////////////////////////////////
// Run with Workers
// ///////////////////////////////////////////////////////////

func Run[IN, OUT any](ctx context.Context, in <-chan IN, transform func(IN) OUT, numWorkers int) <-chan OUT {

	if numWorkers < 1 {
		panic("numWorkers must be >= 1")
	}

	wg := sync.WaitGroup{}

	out := make(chan OUT)

	for i := 0; i < numWorkers; i++ {
		wg.Go(func() {
			for {
				val, ok := CancelOrReceiveBetter(ctx, in)
				if !ok {
					return
				}

				newVal := transform(val)

				if !CancelOrSend(ctx, out, newVal) {
					return
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

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

func Square(in int) int {
	sq := in * in
	fmt.Println("Squared Value: ", in, sq)
	time.Sleep(time.Second)
	return sq
}

func Double(in int) int {
	dubs := in * 2
	fmt.Println("Doubled Value: ", in, dubs)
	return dubs
}

func AddHalf(in int) float64 {
	return float64(in) + 0.5
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	ch1 := Generate(ctx)
	ch2 := Run(ctx, ch1, Square, 10)
	ch3 := Run(ctx, ch2, Double, 1)
	ch4 := Run(ctx, ch3, AddHalf, 1)

	for val := range ch4 {
		fmt.Println("Workflow Output: ", val)
	}
}
