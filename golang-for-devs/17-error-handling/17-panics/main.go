package main

import "fmt"

/////////////////////////////////////////////////////////////
// Panics
/////////////////////////////////////////////////////////////

func boom() {

	defer func() {
		r := recover()

		if r != nil {
			fmt.Println("Recovery Logic")
		}
	}()

	panic("something wrong")
	fmt.Println("Never Executes")
}

func main() {
	boom()
	fmt.Println("Recovered!")
}
