package main

import "fmt"

////////////////////////////////////////////////////////////
// Interface Gotcha
////////////////////////////////////////////////////////////

type MyError struct{}

func (e *MyError) Error() string {
	return "my error message"
}

func DoSomething() error {

	// Good - return nil
	return nil

	// Bad - Returned value is NOT NIL
	// err := &MyError{}
	// return err
}

func main() {
	err := DoSomething()
	if err != nil {
		fmt.Println("Error not nil")
	} else {
		fmt.Println("Error nil")
	}
}
