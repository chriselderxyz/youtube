package main

import (
	"cmp"
	"fmt"
)

////////////////////////////////////////////////////////////
// Generics in Function
////////////////////////////////////////////////////////////

// One generic type
// T -> The generic type, usable in input and result parameters
// Ordered -> The type constraint. T can only be types that support <, > operations
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func CallGenericFunc() {
	// Call it normally
	// Go uses type inference
	x := Max(1, 5)

	// Explicitly set the type of T to [int]
	y := Max[int](2, 48)
}

// More complex example
// Function that takes a slice of any type and transforms
// it into a slice of any other type using a user
// deffined transform function
func MapSlice[T, R any](input []T, fn func(T) R) []R {
	results := make([]R, len(input))

	for index, value := range input {
		results[index] = fn(value)
	}

	return results
}

////////////////////////////////////////////////////////////
// Generics in Structs
////////////////////////////////////////////////////////////

type MysteryBox[T any] struct {
	Value T // The generic type can be used in struct fields
}

// The generic type from the struct definition can be used in methods
func (m MysteryBox[T]) Open() T {
	return m.Value
}

func CreateMysteryBox() {
	m := MysteryBox[int]{Value: 4}
	fmt.Println(m.Value)
}

// Just like functions, you can use multiple generics
type MysteryBoxDeluxe[T1 any, T2 comparable] struct {
	Value1 T1
	Value2 T2
}

////////////////////////////////////////////////////////////
// Generics in Other Types
////////////////////////////////////////////////////////////

// Custom slice type that can hold any type
type List[T any] []T

// For map types, the Key must be comparable
type Dict[K comparable, V any] map[K]V

// Interface type with generics
// Works similar to structs
type Store[T any] interface {
	Save(T) error
	Load() (T, error)
}

////////////////////////////////////////////////////////////
// Custom Constraints
////////////////////////////////////////////////////////////

// Interface type constraint
// Any type with an underlying type of int or string allowed
// Type must implement the Print() method
// NOTE: An interface with a Type Set, can't be used as a regular interface
type StringInt interface {
	~int | ~string
	Print()
}

// Using the custom interface Type Constraint
type BoxBox[T StringInt] struct {
	Value T
}
