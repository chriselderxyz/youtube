package main

import (
	"fmt"
	"sync"
)

//////////////////////////////////////////////////////////////
// WaitGroup - Using wg.Go()
//////////////////////////////////////////////////////////////

func main() {

	wg := sync.WaitGroup{}

	fmt.Println("Main Go Routine")

	wg.Go(Count)

	wg.Wait()
	fmt.Println("Done")
}

func Count() {
	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
	}
}
