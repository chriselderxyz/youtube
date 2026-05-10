package main

import (
	"context"
	"fmt"
)

/////////////////////////////////////////////////////////////
// The Closed Channel Problem
/////////////////////////////////////////////////////////////

func main() {
	for i := 0; i < 20; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		ch := make(chan int)

		close(ch)
		cancel()

		select {
		case <-ctx.Done():
			fmt.Println("Cancel")
		case val := <-ch:
			fmt.Println("Closed", val)
		}
	}
}
