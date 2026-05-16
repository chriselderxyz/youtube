package main

import (
	"fmt"
	"sync"
)

/////////////////////////////////////////////////////////////
// sync.Once - Wrappers
/////////////////////////////////////////////////////////////

func main() {
	f1 := sync.OnceFunc(func() {
		fmt.Println("Runs Once")
	})

	f1()
	f1()
	f1()

	f2 := sync.OnceValue(func() int {
		fmt.Println("Running OnceValue")
		return 10
	})

	a := f2()
	b := f2()
	c := f2()

	fmt.Println(a, b, c)

	f3 := sync.OnceValues(func() (int, string) {
		fmt.Println("Running OnceValues")
		return 10, "ten"
	})

	d, e := f3()
	f, g := f3()
	fmt.Println(d, e, f, g)
}
