package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Context Done Channel
/////////////////////////////////////////////////////////////

func Bark(ctx context.Context, ch chan<- string) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- "Bark":
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func Meow(ctx context.Context, ch chan<- string) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- "Meow":
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	dogs := make(chan string)
	cats := make(chan string)

	wg.Go(func() {
		Bark(ctx, dogs)
	})
	wg.Go(func() {
		Meow(ctx, cats)
	})

	wg.Go(func() {
		for {
			select {
			case <-ctx.Done():
				return
			case val := <-dogs:
				fmt.Println(val)
			case val := <-cats:
				fmt.Println(val)
			}
		}
	})

	wg.Wait()

}
