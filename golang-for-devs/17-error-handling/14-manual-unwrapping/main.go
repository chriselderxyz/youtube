package main

import (
	"errors"
	"fmt"
)

/////////////////////////////////////////////////////////////
// Manual Unwrapping
/////////////////////////////////////////////////////////////

func main() {
	err1 := errors.New("Error 1")
	err2 := errors.New("Error 2")
	err3 := errors.New("Error 3")

	w1 := fmt.Errorf("Wrapped: %w", err1)
	singleUnwrap := errors.Unwrap(w1)
	fmt.Println("Single Unwrap: ", singleUnwrap)

	// Internally
	i, ok := w1.(interface {
		Unwrap() error
	})
	if ok {
		fmt.Println("Assertion Unwrapped: ", i.Unwrap())
	}

	j1 := errors.Join(err1, err2, err3)
	j2 := fmt.Errorf("Joined: %w, %w, %w", err1, err2, err3)

	i2, ok2 := j1.(interface {
		Unwrap() []error
	})
	if ok2 {
		fmt.Println("Manaul Join Unwrap 1: ", i2.Unwrap())
	}

	i3, ok3 := j2.(interface {
		Unwrap() []error
	})
	if ok3 {
		fmt.Println("Manaul Join Unwrap 1: ", i3.Unwrap())
	}
}
