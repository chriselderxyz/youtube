package main

import (
	"errors"
	"fmt"
)

/////////////////////////////////////////////////////////////
// Control Signals
/////////////////////////////////////////////////////////////

var EOF = errors.New("EOF")

func ReadFile() error {
	failed := false

	if failed {
		return errors.New("Failed to Read")
	} else {
		return EOF
	}
}

func main() {
	fmt.Println("Control Signal or Error: ", ReadFile())
}
