package main

import (
	"fmt"
	"sync"
)

/////////////////////////////////////////////////////////////
// sync.RWMutex
/////////////////////////////////////////////////////////////

func main() {
	wg := sync.WaitGroup{}
	counter := 0

	m := sync.RWMutex{}

	add := func() {
		m.Lock()
		counter++
		m.Unlock()
	}

	read := func() {
		m.RLock()
		fmt.Println("Reading Count: ", counter)
		m.RUnlock()
	}

	for i := 0; i < 1000; i++ {
		wg.Go(add)
		wg.Go(read)
	}

	wg.Wait()
	fmt.Println("count: ", counter)
}
