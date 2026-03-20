package main

import "fmt"

////////////////////////////////////////////////////////////
// Basic Function Use
////////////////////////////////////////////////////////////

func Basic() {
	fmt.Println("Do Something")
}

func WithParams(i int, x, y string){
	fmt.Println(i, x, y)
}

func CallingFunction() {
	Basic()

	WithParams(10, "1", "2")

	// Not Allowed
	// Go doesn't let you name or reorder arguments
	// Arguments are always positional, and all args are required
	WithParams(x: "1", i: 10, y: "2") // invalid
}

////////////////////////////////////////////////////////////
// Variadic Parameters
////////////////////////////////////////////////////////////

// nums passed into the function as a slice
func Sum(nums ...int) int {
	total := 0

	for index, num := range nums {
		total += num
	}

	return total
}

func CallVariadic() {
	// 1. Pass multiple values
	sum1 := Sum(1, 2, 3, 4, 5, 6)

	// 2. Use a predefined slice
	nums := []int{1, 2, 3}
	sum2 := Sum(nums...)
}

////////////////////////////////////////////////////////////
// Result Parameters
////////////////////////////////////////////////////////////

// A single, un-named result parameter doesn't need brackets
func SingleResultParam(i int) int {
	return i + 2
}

// Multiple result parameters needs brackets
func MultipleResultParams(i int) (int, int) {
	return i + 2, i - 2
}

func CallFuncWithResultParams() {
	x := OneReturnValue(10)
	y, z := TwoReturnValues(10)

	// Use _ to skip the first return value
	_, z := TwoReturnValues(10)
}

////////////////////////////////////////////////////////////
// Named Result Parameters
////////////////////////////////////////////////////////////

func NamedResultParam() (x, y int) {
	// Named result parameters are initialized to their zero value
	// x = 0, y = 0

	// Naked return
	// Returns 0, 0 (current x, y values)
	return

	// You can reassign the named parameters
	x = 10
	return // 10, 0

	// Regular return still allowed
	// Convention is to not mix regular with naked returns in 1 function
	return 1, 2
}

// Can also name result parameters in interface method definitions
// This is purely for documentation
type Gogo interface {
	MyFunc() (x string, err error)
}

////////////////////////////////////////////////////////////
// Functions as Types
////////////////////////////////////////////////////////////

func FunctionParam(f func(i int) int) func(s string) string {
	return func(s string) string {
		return "Functions are Fun"
	}
}

// Top level functions can be passed to another function
func Double(i int) int {
	return i * 2
}

func CallFunctionWithFunctionParam() {
	
	// Call with an anonymous function
	f1 := FunctionParam(func (i int) int{
		return i + 1
	})

	// Call with a local function defined in this scope
	double := func(i int) int {
		return i * 2
	}
	f2 := FunctionParam(double)

	// Call with a top level function
	f3 := FunctionParam(Double)
}

////////////////////////////////////////////////////////////
// Generics in Functions
////////////////////////////////////////////////////////////

// One generic type T
// T -> The name of the generic type
// Ordered -> The type constraint
func Min[T cmp.Ordered] (a, b T) T {
	if a < b {
        return a
    }
    return b
}

// Multiple generic types
func Pair[A, B any](a A, b B) (A, B) {
	return a, b
}

////////////////////////////////////////////////////////////
// Custom Constraints
////////////////////////////////////////////////////////////

// Generic type N uses Number interface as the type constraint
func Add[N Number](x, y N) N {
	return x+y
}

type Number interface {
	~int | ~float64
}

// The ~ in ~int means derived types work too
type MyType int

////////////////////////////////////////////////////////////
// Calling Functions With Generics
////////////////////////////////////////////////////////////

// Same function from above
func Add2[N Number](x, y N) N {
	return x + y
}

func CallGenerics() {
	// Call it normally
	// Go uses type inference to determine the type for N
	x := Add2(1, 2)

	// Explicitly set type of N to [int]
	y := Add2[int](1, 3)
}

////////////////////////////////////////////////////////////
// Methods
////////////////////////////////////////////////////////////

type MyType struct {}

// (t MyType) -> Attaches the method to the MyType type
// t -> The receiver name - can use in the method
// MyType -> The receiver type
func (t MyType) MyMethod() {
	fmt.Println(t)
}

func CallMethod() {
	m := MyType{}
	m.MyMethod()  // Uses simple dot notation
}

////////////////////////////////////////////////////////////
// Deferring Function Execution
////////////////////////////////////////////////////////////

func Deferring() {
	// defer makes the function call wait until the outer function finishes
	// deferred calls are Last In First Out
	defer fmt.Println("Prints Last")
	defer fmt.Println("Prints Second")
	fmt.Println("Prints First")
}

////////////////////////////////////////////////////////////
// Defer Gotcha
////////////////////////////////////////////////////////////

func DeferGotcha() {
	var e error // nil error

	defer func(errArg error) {

		// Arguments are evaluated immediately
		// So errArg is always nil
		// Even if e is set to a non-nil value later in the function
		if errArg != nil {
			fmt.Println(errArg) // Never executes
		}

		// Accessing outer scope variable (Closure Semantics)
		// This happens when the function runs
		// So e is updated correctly
		if e != nil {
			fmt.Println(e)
		}
	}(e)

	val, ok := TrySomething()
	if !ok {
		// Assign a non-nil error to e if an error occured
		// This only applies to the non-argument error in the deferred function
		e = errors.New("Something Broke")
	}
}

////////////////////////////////////////////////////////////
// Modify Return in Defer
////////////////////////////////////////////////////////////

// Returns 11
func DeferModifiedReturn() (x int) {
	defer func() {
		x ++
	}()

	return 10
}
