package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////
// Creating
////////////////////////////////////////////////////////////

// With string literal
var s string = "hello"

// Zero Value: ""
var s2 string

// From a byte slice
func ReadBytes() string {
	data, _ := os.ReadFile("file.txt")
	return string(data)
}

// With rune literal
var r1 rune = '@'

// Zero Value: 0
var r2 rune

////////////////////////////////////////////////////////////
// Common Operations
////////////////////////////////////////////////////////////

func Length() {
	// Number of bytes in the string
	s1 := "hello"
	l1 := len(s1)
	fmt.Println(l1) // 5

	// Number of bytes - but é takes 2 bytes
	s2 := "héllo"
	l2 := len(s2)
	fmt.Println(l2) // 6

	// Number of runes/characters in the string
	runes := []rune(s2)
	l3 := len(runes)
	fmt.Println(l3) // 5
}

func Compare() {
	s1 := "hello"
	s2 := "world"

	// Equality operator
	equal := s1 == s2

	// strings.Compare function
	// 0 if equal
	// 1 if s1 > s2
	// -1 if s1 < s2
	val := strings.Compare(s1, s2)

	// EqualFold
	equals := strings.EqualFold(s1, s2)

	// Using other comparison operators
	less := s1 < s2
	greater := s1 > s2
	lessOrEqual := s1 <= s2
	greaterOrEqual := s1 >= s2
}

func Concatenate() {
	s1 := "hello"
	s2 := "world"

	// Concatenate with + operator
	// Allocates a new string for each + (inefficient)
	s3 := s1 + " " + s2

	// Concatenate with strings.Builder
	// Efficient - allocates a buffer one time
	// Use this for more than 5 concatenations
	b := strings.Builder{}
	b.Grow(1024)
	b.WriteString(s1)
	b.WriteString(" ")
	b.WriteString(s2)
	b.String()
}

func AccessIndex() {
	s := "hello"
	b1 := s[2] // "l" - returns a byte

	s2 := "héllo"
	b2 := s2[1] // First half é - nonsense byte

	runes := []rune(s2)
	r := runes[1] // 'é' - Actual character
}

func Substring() {

	s := "world"

	// Slicing a string
	// First index is inclusive
	// Second index is exclusive
	sub1 := s[0:3] // "wor"
	sub2 := s[:3]  // "wor"
	sub3 := s[3:]  // "ld"

	// Slices use bytes, not runes
	s2 := "héllo"
	sub4 := s2[0:3] // expect "hél", get "hé"
}

func StringConversion() {

	// Integer to a string
	s := strconv.Itoa(123)

	// String to an integer
	i, err := strconv.Atoi("123")
}

func StringIteration() {
	s := "héllo"

	// Iterate over the BYTES in the string
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// Iterate over the RUNES in the string
	for index, rune := range s {
		fmt.Println(index, rune)
	}
}
