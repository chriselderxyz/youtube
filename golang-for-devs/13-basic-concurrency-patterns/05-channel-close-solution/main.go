package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Closing Channel Solution
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

	ctxLong, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ctxShort, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	dogs := make(chan string)
	cats := make(chan string)

	wg.Go(func() {
		Bark(ctxShort, dogs)
	})
	wg.Go(func() {
		Meow(ctxLong, cats)
	})

	wg.Go(func() {
		for {
			select {
			case <-ctxLong.Done():
				return
			case val, ok := <-dogs:
				if !ok {
					dogs = nil
					break
				}
				fmt.Println("Dogs: ", val)
			case val, ok := <-cats:
				if !ok {
					cats = nil
					break
				}
				fmt.Println("Cats: ", val)
			}
		}
	})

	wg.Wait()

}
