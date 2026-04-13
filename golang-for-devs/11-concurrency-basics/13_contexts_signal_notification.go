package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

////////////////////////////////////////////////////////////
// Contexts - signal.NotifyContext
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
	ctx2, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	ctxWithValue := context.WithValue(ctx2, requestIDKey, "12345")

	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Go(func() {
		Generate(ctxWithValue, ch)
	})

	wg.Go(func() {
		Transform(ch)
	})

	wg.Wait()
}
