package main

import (
	"fmt"
	"sync"
)

/////////////////////////////////////////////////////////////
// sync.Mutex
/////////////////////////////////////////////////////////////

func main() {
	wg := sync.WaitGroup{}
	counter := 0

	m := sync.Mutex{}

	add := func() {
		m.Lock()
		counter++
		m.Unlock()
	}

	for i := 0; i < 1000; i++ {
		wg.Go(add)
	}

	wg.Wait()
	fmt.Println("count: ", counter)
}
