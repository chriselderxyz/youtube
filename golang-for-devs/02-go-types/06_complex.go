package main

////////// Complex Numbers //////////

// 64-bit complex number
// 1 float32 for real part
// 1 float32 for imaginary part
var c64 complex64 = 1 + 2i

// 128-bit complex number
// 1 float64 for real part
// 1 float64 for imaginary part
var c128 complex128 = 1 + 2i

// Zero Value - 0 + 0i for all complex types
var cZero complex64

// Type Inference
// Go assumes a complex128 by default
var cInfer = 1 + 2i
