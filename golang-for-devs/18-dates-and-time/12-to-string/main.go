package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// Time to String
/////////////////////////////////////////////////////////////

func main() {
	t := time.Now()

	str1 := t.String()
	str2 := t.Format("Jan 2, 2006")
	fmt.Println("t.String() - ", str1)
	fmt.Println("t.Format() - ", str2)

	// AppendText
	// AppendFormat
}
