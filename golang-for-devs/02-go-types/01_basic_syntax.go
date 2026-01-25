package main

// Package-Level Variables
// Accessible throughout the package
var packageLevelVar string = "package-level"

// Public Variables
// Accessible outside the package
var PublicVar string = "public"

func BasicSyntax() {

	// Local Variables
	var localVar string = "world"
	var noType = "var"

	// Short hand with Type Inference
	// ONLY works in functions
	shortHand := "hello"

	// No value specified
	// Initialized with the Zero Value for the type
	var zeroValue int // 0

	// Casting
	var old int8 = 100
	new := int(old)
}
