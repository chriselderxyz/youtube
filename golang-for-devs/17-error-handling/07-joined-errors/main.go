package main

import (
	"errors"
	"fmt"
)

/////////////////////////////////////////////////////////////
// Joined Errors
/////////////////////////////////////////////////////////////

func main() {
	err1 := DoSomething()
	err2 := DoSomethingElse()

	j := errors.Join(err1, err2) // Logs with \n's
	j2 := fmt.Errorf("error 1: %w, error 2: %w", err1, err2) // Logs without \n's

	fmt.Println("j: ", j)
	fmt.Println("j2: ", j2)
}

func DoSomething() error {
	return errors.New("do something error")
}

func DoSomethingElse() error {
	return errors.New("do something else error")
}
