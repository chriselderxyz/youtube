package main

import "fmt"

/////////////////////////////////////////////////////////////
// Channels - Unbuffered Channels
/////////////////////////////////////////////////////////////

func main() {
	ch := make(chan int)

	// Sending to, or Receiving from, a channel blocks the Go Routines execution
	// Deadlock.
	ch <- 10
	val := <-ch

	fmt.Println(val)
}
