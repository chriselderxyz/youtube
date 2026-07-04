package variables

import "sync"

/////////////////////////////////////////////////////////////
// The Many Ways to Create Variables in Go
/////////////////////////////////////////////////////////////

// Top level variables must use var
var x = 10

// const's can be inside or outside a function
const outside = 11

func main() {

	/////////////////////////////////////////////////////////////
	// var
	/////////////////////////////////////////////////////////////
	var x1 int = 10 // Full form
	var x2 = 10     // Type inference
	var x3 int      // Zero value instantiation
	var x4 int = 0  // Equivalent to zero value instantiation

	var s1 string // ""    - Zero value instantiation
	var b1 bool   // false - Zero value instantiation

	var p *int       // nil - Zero value instantiation
	var slices []int // nil - Zero value instantiation
	var arr [3]int   // [0, 0, 0] - Zero value instantiation

	var wg sync.WaitGroup   // Empty WG - Zero value instantiation
	var wgp *sync.WaitGroup // nil - Zero value instantiation

	/////////////////////////////////////////////////////////////
	// Shorthand
	/////////////////////////////////////////////////////////////
	var z1 int = 10   // Longform
	z2 := 10          // int   // Shortform with type inference (int)
	z3 := int64(10)   // Shortform with type conversion (int64)
	var z4 int64 = 10 // Longform with specific type

	/////////////////////////////////////////////////////////////
	// make
	/////////////////////////////////////////////////////////////
	var slice2 []int = []int{}    // Empty slice (len, cap = 0)
	notEmpty := make([]int, 3, 6) // Non-empty slice - [0, 0, 0] (len = 3, cap = 6)

	m1 := make(map[string]int, 10) // Empty map with size hint
	m2 := map[string]int{}         // Empty map with no size hint

	ch1 := make(chan int, 100) // Usable channel
	var ch2 chan int           // nil - non-usable channel

	/////////////////////////////////////////////////////////////
	// new
	/////////////////////////////////////////////////////////////
	n := new(int)         // non-nil pointer to an int set to 0
	pointer := new([]int) // non-nil pointer to a NIL slice - Tricky

	/////////////////////////////////////////////////////////////
	// const
	/////////////////////////////////////////////////////////////
	const c = 10
}
