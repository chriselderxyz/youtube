package main

import (
	"fmt"
	"sync"
)

//////////////////////////////////////////////////////////////
// WaitGroup - Basics
//////////////////////////////////////////////////////////////

func main() {

	wg := sync.WaitGroup{}

	fmt.Println("Main Go Routine")

	wg.Add(1)
	go Count(&wg)

	wg.Wait()
	fmt.Println("Done")
}

func Count(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
	}
}
