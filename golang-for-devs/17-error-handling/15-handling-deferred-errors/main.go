package main

import (
	"errors"
	"fmt"
)

/////////////////////////////////////////////////////////////
// Handling Deferred Errors
/////////////////////////////////////////////////////////////

func ReturnsErr() error {
	return errors.New("New Error")
}

func DoSomething() (err error) {
	fmt.Println("Do Something")
	defer func() {
		e := ReturnsErr()
		if e != nil {
			err = e
		}
	}()

	return
}

func main() {
	e := DoSomething()
	fmt.Println(e)
}
