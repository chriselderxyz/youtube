package main

import (
	"fmt"
	"sync"
)

/////////////////////////////////////////////////////////////
// sync.Map
/////////////////////////////////////////////////////////////

func main() {
	m := sync.Map{}

	m.Store("key", "value")

	v1, _ := m.Load("key")
	fmt.Println("Value: ", v1)

	v2, ok := m.Load("Not Found")
	if !ok {
		fmt.Println("Not Found: ", v2)
	}

	m.Delete("key")
	m.Clear()

	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")

	m.Range(func(key any, val any) bool {
		fmt.Println(key, val)
		return true
	})
}
