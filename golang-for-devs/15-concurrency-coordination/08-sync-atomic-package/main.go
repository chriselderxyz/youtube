package main

import (
	"fmt"
	"sync/atomic"
)

/////////////////////////////////////////////////////////////
// sync/atomic Package
/////////////////////////////////////////////////////////////

func main() {
	b := atomic.Bool{}

	b.Store(true)

	v1 := b.Load()

	v2 := b.Swap(false)
	fmt.Println(v1, v2)

	b.CompareAndSwap(true, false)

	a := int64(10)
	atomic.AddInt64(&a, 100)
}
