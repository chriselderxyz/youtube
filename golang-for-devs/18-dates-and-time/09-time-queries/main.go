package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Time - queries
/////////////////////////////////////////////////////////////

func main() {
	t := time.Now()

	isDST := t.IsDST()

	start, end := t.ZoneBounds()

	if t.IsZero() {
		fmt.Println("Not Set")
	}

	fmt.Println("t := time.Now(): ", t)
	fmt.Println("t.IsDST(): ", isDST)
	fmt.Println("t.ZoneBounds() -> start: ", start, " end: ", end)
}
