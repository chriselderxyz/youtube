package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////
// Channels - Worker Pattern
/////////////////////////////////////////////////////////////

func Generate(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 20; i++ {
		ch <- rand.IntN(100) + 1
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

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Go(func() {
		Generate(ch)
	})

	// Spin up 10 "worker" Go Routines running the Transform function
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			Transform(ch)
		})
	}

	wg.Wait()
}









































// func Generate(ch chan<- int) {
// 	for i := 0; i < 20; i++ {
// 		ch <- rand.IntN(100) + 1
// 	}

// 	close(ch)
// }

// // One

// func Transform(ch chan int) {
// 	for val := range ch {
// 		val = val * 10
// 		fmt.Println(val)
// 	}
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	ch := make(chan int)

// 	wg.Go(func() {
// 		Generate(ch)
// 	})

// 	wg.Go(func() {
// 		Transform(ch, 1)
// 	})

// 	wg.Wait()
// }

// // Two
// func main() {
// 	wg := sync.WaitGroup{}
// 	ch := make(chan int)

// 	wg.Go(func() {
// 		Generate(ch)
// 	})

// 	for i := 0; i < 10; i++ {
// 		wg.Go(func() {
// 			Transform(ch, i)
// 		})
// 	}

// 	wg.Wait()
// }

// func Transform(ch <-chan int, i int) {
// 	for val := range ch {
// 		val = val * 10
		// time.Sleep(time.Duration(rand.IntN(2000)+1) * time.Millisecond)
// 		fmt.Println(val, "Go Routine: ", i)
// 	}
// }
