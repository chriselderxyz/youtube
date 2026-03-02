package main

////////////////////////////////////////////////////////////
// Working with Pointers
////////////////////////////////////////////////////////////

func CreatingPointers(){
	var i int = 10
	var p *int = &i    // Pointer to i
	var pp **int = &p  // Pointer to p (pointer to a pointer)

	p2 := new(int) // Pointer to a 'new' int - 0

	// Zero Value = nil
	var p3 *int
}

func GetValueFromPointer() {
	i := 10
	p := &i

	value := *p // Dereferencing
}

func NonAddressables(){ 
	const x = 10
	p1 := &x // error

	p2 := &10 // error
	p3 := &"hello" // error

	p4 := &(x + y) // error

	func f() int { return 10 }
	p5 := &f() // error

	i := 10
	p6 := &float64(i) // error

	// Fix with an Intermediary Variable
	result := f()
	p7 := &result // Works
}

////////////////////////////////////////////////////////////
// Using Pointers with Functions
////////////////////////////////////////////////////////////

func MyFunc(i *int) {}

func MyOtherFunc(){
	i := 10
	MyFunc(&i)
}

////////////////////////////////////////////////////////////
// Heap Escaping Gotcha
////////////////////////////////////////////////////////////

// Bad Performance
// The returned pointer references a local variable to the function
// If the variable stays on the stack, it may get garbage collected
// So the compiler escapes i to the heap
func ShareUp() *int {
	i := 10
	return &i
}

// Good Performance
func ShareDown(i *int) {
	*i = 10
}
