package main

import (
	"errors"
	"fmt"
)

/////////////////////////////////////////////////////////////
// Wrapped Errors
/////////////////////////////////////////////////////////////

func main() {
	wrapped := WrappedError()
	mainErr := fmt.Errorf("main error: %w", wrapped)
	fmt.Println(mainErr)
}

func WrappedError() error {
	nested := ReturnsError()
	outer := fmt.Errorf("outer failed: %w", nested)
	return outer
}

func ReturnsError() error {
	return errors.New("deepest error")
}
