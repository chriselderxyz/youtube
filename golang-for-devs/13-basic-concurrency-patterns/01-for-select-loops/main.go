package main

import (
	"fmt"
	"sync"
	"time"
)

////////////////////////////////////////////////////////////
// For Select Loops
////////////////////////////////////////////////////////////

func Bark(ch chan<- string) {
	defer close(ch)
	for {
		ch <- "Bark"
		time.Sleep(500 * time.Millisecond)
	}
}

func Meow(ch chan<- string) {
	defer close(ch)
	for {
		ch <- "Meow"
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	wg := sync.WaitGroup{}
	dogs := make(chan string)
	cats := make(chan string)

	wg.Go(func() {
		Bark(dogs)
	})
	wg.Go(func() {
		Meow(cats)
	})

	wg.Go(func() {
		for {
			select {
			case val := <-dogs:
				fmt.Println(val)
			case val := <-cats:
				fmt.Println(val)
			}
		}
	})

	wg.Wait()
}
