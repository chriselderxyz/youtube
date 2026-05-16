package main

import (
	"fmt"
	"sync"
)

/////////////////////////////////////////////////////////////
// sync.Once
/////////////////////////////////////////////////////////////

type Store struct {
	number int
}

var store *Store
var once sync.Once

func GetStore(i int) *Store {
	once.Do(func() {
		fmt.Println("Opening Connection")
		store = &Store{number: i}
	})

	return store
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			s := GetStore(i)
			fmt.Println("Accessing Store: ", s)
		})
	}

	wg.Wait()
}
