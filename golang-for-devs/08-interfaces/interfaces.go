package main

import "fmt"

////////////////////////////////////////////////////////////
// Defining Interfaces
////////////////////////////////////////////////////////////

type Printer interface {
	Print(message string)
}

type FancyPrinter interface {
	Printer // Embedding interfaces
	PrintFancy(message string)
}

////////////////////////////////////////////////////////////
// Implementing an Interface
////////////////////////////////////////////////////////////

// With a struct type
type FaxMachine struct {
	model  string
	number int
}

func (f FaxMachine) Print(m string) {
	fmt.Println(m)
}

// With a custom int type
type MyInt int

func (i MyInt) Print(m string) {
	fmt.Println(i, m)
}

////////////////////////////////////////////////////////////
// Receivers Gotcha
////////////////////////////////////////////////////////////

type Dog interface {
	Bark()
}

// A struct type that implements the Dog interface
type Sparky struct{}

// Sparky implementing Dog using a Value Receiver
// func (s Sparky) Bark() {
// 	fmt.Println("Bark!")
// }

// Sparky implementing Dog using a Pointer Receiver
func (s *Sparky) Bark() {
	fmt.Println("Bark!")
}

func GoDogGo() {

	// A value instance is only a "Dog" if
	// the Bark method uses a Value Receiver
	val := Sparky{}

	// A pointer instance is always a "Dog"
	// It doesn't matter what receiver type is used
	pointer := &Sparky{}

	// val is only a Dog if a Value Receiver is used
	// Passing a "non-Dog" to Speak throws a compile error
	// If a pointer receiver is used, this is invalid
	Speak(val)

	// Always Works
	Speak(pointer)

	// This works
	// Even if Bark uses a Pointer Receiver
	val.Bark()
}

func Speak(d Dog) {
	d.Bark()
}

////////////////////////////////////////////////////////////
// Receivers Gotcha With Embedded Structs
////////////////////////////////////////////////////////////

type I interface {
	Foo()
}

type Inner struct{}

func (i *Inner) Foo() {}

// Same rules apply
// Only a pointer Outer instance satisfies the I interface
// because a pointer receiver was used to implement I on Inner
type Outer struct {
	Inner
}

////////////////////////////////////////////////////////////
// Using Interface Types, in General
////////////////////////////////////////////////////////////

// Interface
type CatInter interface {
	Meow()
}

// Type implementing the CatInter interface
type CatStruct struct{}

func (a CatStruct) Meow() {
	fmt.Println("Meow")
}

// Function needs to return a CatInter instance
// How do we make one of those?
func GoCatGo() CatInter {
	// Implicitly
	// Go auto-converts the CatStruct to CatInter when you return
	catStruct := CatStruct{}
	return catStruct

	// With Explicit Type Conversion
	// Type conversion works because CatStruct implements CatInter
	catInterface1 := CatInter(CatStruct{})
	return catInterface1

	// With an Explicit Type using var
	var catInterface2 CatInter = CatStruct{}
	return catInterface2
}

////////////////////////////////////////////////////////////
// Interface -> Type Conversion
////////////////////////////////////////////////////////////

func InterfaceToStruct() {
	catInterface := CatInter(CatStruct{})

	// Compiler Error
	// You can't use regular type conversion to go from Interface to Type
	structInstance := CatStruct(catInterface)

	// You need to use runtime type assertion instead
	structInstance, ok := catInterface.(CatStruct) // works, may panic
	if !ok {
		// handle wrong type
		fmt.Println(structInstance)
	}
}

////////////////////////////////////////////////////////////
// Empty Interfaces
////////////////////////////////////////////////////////////

// Old way of having a generic argument type
func Empty(t interface{}) {
	fmt.Println(t)
}

// New generics syntax (same same)
func WithGenerics(t any) {
	fmt.Println(t)
}

////////////////////////////////////////////////////////////
// nil Error Gotcha
////////////////////////////////////////////////////////////

// error interface built into Go:
// type error interface {
// 	Error() string
// }

// Custom error type that implements error
type MyError struct{}

func (e *MyError) Error() string { return "boom" }

func f() error {

	// No error occurs
	// I try to use a nil pointer to MyError type
	var e *MyError = nil

	// Go converts to a non-nil interface containing a nil pointer
	// Looks like an error occured when there was none
	return e

	// What you should actually do if there is no error
	return nil
}
