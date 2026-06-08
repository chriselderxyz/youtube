package main

import "fmt"

/////////////////////////////////////////////////////////////
// Custom Error Types
/////////////////////////////////////////////////////////////

type MyError struct {
	Code int
	UserID string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("my error, status code: %d, user: %s", e.Code, e.UserID)
}

func main() {
	err := &MyError{
		Code: 500,
		UserID: "12345",
	}

	fmt.Println(err)
}
