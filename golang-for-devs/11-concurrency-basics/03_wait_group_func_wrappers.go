package main

import (
	"fmt"
	"sync"
)

//////////////////////////////////////////////////////////////
// WaitGroup - Using Anonymous Functions
//////////////////////////////////////////////////////////////

func main() {

	wg := sync.WaitGroup{}

	fmt.Println("Main Go Routine")

	wg.Add(1)
	go func() {
		Count()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done")
}

func Count() {
	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
	}
}
