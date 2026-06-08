package main

import (
	"errors"
	"fmt"
)

/////////////////////////////////////////////////////////////
// Sentinel Errors
/////////////////////////////////////////////////////////////

var ErrNotFound = errors.New("Not Found")

func FindUser(id string) error {
	// Check for user
	// Return error if not found
	return ErrNotFound
}

func main() {
	err := FindUser("id")
	fmt.Println("Error: ", err)
}
