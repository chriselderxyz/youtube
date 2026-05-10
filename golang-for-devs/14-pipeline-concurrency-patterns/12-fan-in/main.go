package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Merge (Fan-In)
/////////////////////////////////////////////////////////////

func Merge[T any](ctx context.Context, buf int, inputs ...<-chan T) <-chan T {
	out := make(chan T, buf)

	wg := sync.WaitGroup{}

	for _, ch := range inputs {
		wg.Go(func() {
			if ch == nil {
				return
			}

			for {
				val, ok := CancelOrReceiveBetter(ctx, ch)
				if !ok {
					return
				}

				if !CancelOrSend(ctx, out, val) {
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

func Split[T any](ctx context.Context, in <-chan T, numChannels int, buf int) []<-chan T {

	if numChannels < 1 {
		panic("numChannels must be >= 1")
	}
	if buf < 0 {
		buf = 0
	}

	chs := make([]chan T, numChannels)
	for i := range chs {
		chs[i] = make(chan T, buf)
	}

	go func() {
		defer func() {
			for _, ch := range chs {
				close(ch)
			}
		}()

		for {
			val, ok := CancelOrReceiveBetter(ctx, in)
			if !ok {
				return
			}

			for _, ch := range chs {
				if !CancelOrSend(ctx, ch, val) {
					return
				}
			}
		}
	}()

	outs := make([]<-chan T, numChannels)
	for i, ch := range chs {
		outs[i] = ch
	}

	return outs
}

func Run[IN, OUT any](ctx context.Context, in <-chan IN, transform func(IN) OUT, numWorkers int, buf int) <-chan OUT {

	if numWorkers < 1 {
		panic("numWorkers must be >= 1")
	}

	wg := sync.WaitGroup{}

	out := make(chan OUT, buf)

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

func AddTen(in int) int {
	return in + 10
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	ch1 := Generate(ctx)
	ch2 := Run(ctx, ch1, Square, 10, 10)

	// Split Pipeline
	forks := Split(ctx, ch2, 2, 10)

	chA1 := Run(ctx, forks[0], Double, 1, 0)
	chB1 := Run(ctx, forks[1], AddTen, 1, 0)

	merged := Merge(ctx, 0, chA1, chB1)

	for val := range merged {
		fmt.Println("output 2: ", val)
	}
}
