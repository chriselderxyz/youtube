package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

////////////////////////////////////////////////////////////
// Contexts - WithCancel & WithValue
////////////////////////////////////////////////////////////

func Generate(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cleanup for Request: ", ctx.Value(requestIDKey))
			close(ch)
			return
		case ch <- rand.IntN(100) + 1:
		}
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

type RequestIDKeyType struct{}

var requestIDKey RequestIDKeyType

func main() {

	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	ctxWithValue := context.WithValue(ctxWithCancel, requestIDKey, "12345")

	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Go(func() {
		Generate(ctxWithValue, ch)
	})

	wg.Go(func() {
		Transform(ch)
	})

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Cancelling Context")
		cancel()
	})

	wg.Wait()
}
