package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Time - Comparison
/////////////////////////////////////////////////////////////

func main() {
	smaller := time.Now()
	larger := smaller.Add(time.Hour * 48)

	before := smaller.Before(larger) // true
	after := larger.After(smaller)   // true
	equal := smaller.Equal(larger)   // false

	// -1 smaller < larger - this
	// 0 smaller = larger
	// 1 smaller > larger
	comp := smaller.Compare(larger)

	// Don't use ==
	eq := smaller == smaller.UTC() // false

	fmt.Println("smaller.Before(larger): ", before)
	fmt.Println("larger.After(smaller): ", after)
	fmt.Println("smaller.Equal(larger): ", equal)
	fmt.Println("smaller.Compare(larger): ", comp)
	fmt.Println("(DONT DO) - smaller == smaller.UTC()", eq)

}
