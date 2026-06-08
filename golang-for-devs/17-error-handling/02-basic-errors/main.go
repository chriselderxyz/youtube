package main

import (
	"errors"
	"fmt"
)

/////////////////////////////////////////////////////////////
// Creating Errors
/////////////////////////////////////////////////////////////

func main() {
	err1 := errors.New("error message")
	err2 := fmt.Errorf("error 2 message: %s", "Current User ID")

	fmt.Println("err1: ", err1)
	fmt.Println("err2: ", err2)
}
