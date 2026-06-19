package main

import (
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// time.Month
/////////////////////////////////////////////////////////////

func main() {
	m1 := time.April
	m2 := time.May

	m3 := time.Month(7) // July

	str := m3.String() // "July"

	fmt.Println("m1 - time.April: ", m1)
	fmt.Println("m2 - time.May: ", m2)
	fmt.Println("m3 - time.Month(7): ", m3)
	fmt.Println("m3 string - m3.String: ", str)
}
