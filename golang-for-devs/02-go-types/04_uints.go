package main

////////// Unsigned Integer Types //////////

// 8-bit unsigned integer
var u8 uint8 = 255

// 16-bit unsigned integer
var u16 uint16 = 65535

// 32-bit unsigned integer
var u32 uint32 = 4294967295
var u64 uint64 = 18446744073709551615

// 32 or 64 bits depending on the architecture
// Best practice to use uint for most cases
// - Slight performance win
// - More idiomatic
var u uint = 10

// Zero Value - 0 for all unsigned integer types
var uZero uint

// Type Inference
// Go assumes a uint by default
var uInfer = 10

////////// Special Cases //////////

// Alias for uint8
// Represents a byte as a base 10 number
var b byte = 255

// Alais for uint
// Used to hold pointer addresses
// Only used in low-level programming
var uptr uintptr = 0
