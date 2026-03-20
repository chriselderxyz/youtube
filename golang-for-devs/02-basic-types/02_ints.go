package main

////////// Signed Integer Types //////////

// 8-bit signed integer
var i8 int8 = 127

// 16-bit signed integer
var i16 int16 = 32767

// 32-bit signed integer
var i32 int32 = 2147483647

// 64-bit signed integer
var i64 int64 = 9223372036854775807

// 32 or 64 bits depending on the architecture
// Best practice to use int for most cases
// - Slight performance win
// - More idiomatic
var i int = 10

// Zero Value - 0 for all integer types
var iZero int

// Type Inference
// Go assumes an int by default
var iInfer = 10

////////// Special Case //////////

// Alias for int32
// Same as a Character in other languages
// 32-bits (4 bytes) to accommodate all UTF-8 characters 
var r rune = 123
