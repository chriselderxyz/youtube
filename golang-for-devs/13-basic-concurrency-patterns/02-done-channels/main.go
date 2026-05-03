package main

import (
	"fmt"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Done Channels (replaced with contexts)
/////////////////////////////////////////////////////////////

func Bark(done <-chan struct{}, ch chan<- string) {
	defer close(ch)
	for {
		select {
		case <-done:
			return
		case ch <- "Bark":
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func Meow(done <-chan struct{}, ch chan<- string) {
	defer close(ch)
	for {
		select {
		case <-done:
			return
		case ch <- "Meow":
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {

	done := make(chan struct{})

	time.AfterFunc(3*time.Second, func() {
		close(done)
	})

	wg := sync.WaitGroup{}
	dogs := make(chan string)
	cats := make(chan string)

	wg.Go(func() {
		Bark(done, dogs)
	})
	wg.Go(func() {
		Meow(done, cats)
	})

	wg.Go(func() {
		for {
			select {
			case <-done:
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
