package main

import (
	"fmt"
	"sync"
)

/////////////////////////////////////////////////////////////
// sync.Mutex - TryLock
/////////////////////////////////////////////////////////////

func main() {
	wg := sync.WaitGroup{}
	counter := 0

	m := sync.Mutex{}

	add := func() {
		locked := m.TryLock()

		if locked {
			counter++
			m.Unlock()
		}
	}

	for i := 0; i < 1000; i++ {
		wg.Go(add)
	}

	wg.Wait()
	fmt.Println("count: ", counter)
}
