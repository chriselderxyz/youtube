package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

/////////////////////////////////////////////////////////////
// Using == to Check Errors
/////////////////////////////////////////////////////////////

var ErrSentinel = errors.New("sentinel")

type CustomError struct {
	userID string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("CustomError for User %s", e.userID)
}

func DoSomething() error {
	i := rand.IntN(2)
	if i == 0 {
		return ErrSentinel
	} else {
		return &CustomError{}
	}
}

func main() {
	for i := 0; i < 20; i++ {
		err := DoSomething()
		fmt.Println("==============================")
		fmt.Println("ERROR: ", err)
		fmt.Println("==============================")

		fmt.Println("== check Sentinel Error: ", err == ErrSentinel)
		fmt.Println("== check custom error: ", err == &CustomError{})
	}
}
