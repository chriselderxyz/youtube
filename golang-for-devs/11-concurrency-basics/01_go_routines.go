package main

import "fmt"

/////////////////////////////////////////////////////////////
// Go Routines
/////////////////////////////////////////////////////////////

func main() {
	fmt.Println("Main Go Routine")
	go Count()
	fmt.Println("Done")
}

func Count() {
	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
	}
}
