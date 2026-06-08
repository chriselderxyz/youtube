package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

/////////////////////////////////////////////////////////////
// Handling Basic Errors
/////////////////////////////////////////////////////////////

func main() {
	for i := 0; i < 20; i++ {
		err1 := BasicError()

		if err1 != nil {
			// handle error
			fmt.Println("ERROR: ", err1)
		} else {
			fmt.Println("No Error")
		}

		// if err1.Error() == "Basic Error" {
		// 	// Don't do this
		// }
	}
}

func BasicError() error {
	time.Sleep(200 * time.Millisecond)
	i := rand.IntN(2)
	if i == 1 {
		return errors.New("Basic Error")
	} else {
		return nil
	}
}
