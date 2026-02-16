package main

import (
	"fmt"
	"slices"
)

////////////////////////////////////////////////////////////
// Working With Lists
////////////////////////////////////////////////////////////

func Create() {

	// With var syntax
	var a1 [3]int = [3]int{1, 2, 3} // Array
	var s1 []int = []int{1, 2, 3}   // Slice

	// With short hand syntax
	a2 := [3]int{1, 2, 3} // Array
	s2 := []int{1, 2, 3}  // Slice

	// Using make
	// length = 3, capacity = 6
	s3 := make([]int, 3, 6)

	// Zero Value
	var a4 [3]int // [0,0,0] - Array
	var s4 []int  // nil - Slice

	// // Empty slice != nil slice
	s5 := []int{}
}

func AccessAndModify() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}

	// Access
	v1 := a1[0]
	v2 := s1[1]

	// Modify
	a1[1] = 10
	s1[2] = 20
}

func GetLength() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}

	// Using len function
	l1 := len(a1) // 3
	l2 := len(s1) // 3

	// Nil slice has length 0
	var nilSlice []int
	l3 := len(nilSlice) // 0
}

func Iterate() {
	s := []int{1, 2, 3}

	// Loop through the slice
	for index, value := range s {
		fmt.Println(index, value)
	}
}

func Append() {
	s1 := []int{1, 2, 3}

	// Add element to end of the slice
	// If capacity isn't enough, a new backing array is allocated
	s2 := append(s1, 4)
}

func Copying(s []int) {

	// Create a new slice with same length
	newSlice := make([]int, len(s))

	// Copy elements from the source slice to the new slice
	numElements := copy(newSlice, s)
}

////////////////////////////////////////////////////////////
// Slice Append Performance Considerations
////////////////////////////////////////////////////////////

func BadPerformance() {
	slice := []int{}
	for i := 0; i < 100; i++ {
		// Go needs to allocate a new backing array multiple times
		slice = append(slice, i)
	}
}

func GoodPerformance() {
	// Using Make to pre-allocate the backing array
	slice := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		// No reallocation required
		slice = append(slice, i)
	}
}

////////////////////////////////////////////////////////////
// Slicing
////////////////////////////////////////////////////////////

func Slicing() {
	s1 := []int{1, 2, 3, 4, 5}

	// Regular slicing
	// First index is inclusive
	// Second index is exclusive
	s2 := s1[1:3] // Index 1 to 2 - [2, 3]

	// Drop numbers
	s3 := s1[:3] // Index 0 to 2 - [1, 2, 3]
	s4 := s1[3:] // Index 3 to end - [4, 5]

	// Limit capacity with third number
	s5 := s1[0:3:3] // [1, 2] - cap = 3, instead of 5
}

////////////////////////////////////////////////////////////
// Standard Library Operations
////////////////////////////////////////////////////////////

func StandardLibrary() {
	s1 := []int{1, 2, 3}
	s2 := []int{3, 4, 5}

	// Clone
	s3 := slices.Clone(s1)

	// Concatenate
	s4 := slices.Concat(s1, s2)

	// Deleting
	// Works by left shifting elements
	// First index is inclusive, second index is exclusive
	s5 := slices.Delete(s1, 0, 2) // index 0, 1

	// Filtering
	// Works by left shifting elements
	s6 := slices.DeleteFunc(s2, func(val int) bool {
		return val > 4 // Deletes elements that return true
	})

	// Inserting
	// Insert the value 4, at index 1
	// Works by right shifting elements
	// May allocate a new backing array if capacity isn't enough
	s7 := slices.Insert(s1, 1, 4)

	// Sorting - Basic
	// Uses < and > operators to compare elements
	// May swap elements that are equal
	slices.Sort(s1)

	// Sorting - Stable
	// Uses a custom function to compare elements
	// Doesn't change the order of elements that are equal
	slices.SortStableFunc(s1, func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	// Sorting - Check if Slice is Sorted
	sorted := slices.IsSorted(s1) // true

	// Comparisons - Check if Equal
	equal := slices.Equal(s1, s2)

	// Comparisons - More Detailed
	// 1 - s1 > s2
	// 0 - s1 == s2
	// -1 - s1 < s2
	val := slices.Compare(s1, s2)

	// Searching
	index := slices.Index(s1, 2)  // Find first occurrence
	has := slices.Contains(s1, 2) // Check if value is in slice
}
