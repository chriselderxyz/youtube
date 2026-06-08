package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

/////////////////////////////////////////////////////////////
// errors.Is
/////////////////////////////////////////////////////////////

var ErrSentinel = errors.New("Sentinel Error")

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
		return &CustomError{userID: "12345"}
	}
}

func main() {
	for i := 0; i < 20; i++ {

		err := DoSomething()
		fmt.Println("==============================")
		fmt.Println("ERROR: ", err)
		fmt.Println("==============================")

		// == Check
		fmt.Println("== check Sentinel Error: ", err == ErrSentinel)
		fmt.Println("== check custom error: ", err == &CustomError{})

		// errors.Is Checks
		fmt.Println("err Is ErrSentinel: ", errors.Is(err, ErrSentinel))
		fmt.Println("err Is CustomError: ", errors.Is(err, &CustomError{})) // Doesn't work.

		// Works for wrapped and joined errors
		wrapped := fmt.Errorf("Wrapped: %w", err)
		joined := errors.Join(errors.New("New Error"), err)

		fmt.Println("wrapped Is ErrSentinel: ", errors.Is(wrapped, ErrSentinel)) // Works!
		fmt.Println("joined Is ErrSentinel: ", errors.Is(joined, ErrSentinel))
	}
}
