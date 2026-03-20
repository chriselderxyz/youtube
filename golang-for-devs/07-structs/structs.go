package main

import "fmt"

////////////////////////////////////////////////////////////
// Defining a Struct Type
////////////////////////////////////////////////////////////

type Person struct {
	Name    string
	Age     int
	private int // Only accessible inside this package
}

////////////////////////////////////////////////////////////
// Working With Structs
////////////////////////////////////////////////////////////

func CreatingStructs() {

	// Full var syntax
	var p Person = Person{
		Name: "Timmy",
		Age:  10,
	}

	// Short hand syntax
	p2 := Person{Name: "Timmy", Age: 10}

	// Inline pointer syntax
	p3 := &Person{Name: "Timmy", Age: 10}

	// Relying on field ordering (not ideal)
	p4 := Person{"Timmy", 10}

	// Partially populated
	p5 := Person{Name: "Timmy"} // Age 0

	// Zero Value
	// Struct values can't be nil
	var p6 Person  // Name = "", Age 0
	p7 := Person{} // Name = "", Age 0 - Zero Value = Empty Value

	// Use a pointer if you need nil semantics
	var p8 *Person // nil
}

func AccessModifyStructFields() {
	p1 := Person{
		Name: "Timmy",
		Age:  12,
	}

	// Access
	name := p1.Name

	// Modify
	p1.Name = "John"
}

func CopyStructs() {
	a := MyStruct{List: []int{1, 2, 3}}
	b := a // Shallow Copy

	b.List[2] = 100 // Modifies copy and original
}

////////////////////////////////////////////////////////////
// Methods
////////////////////////////////////////////////////////////

type Person struct {
	Name string
	Age  int
}

// Value receiver method
// Can't mutate original caller
func (p Person) Birthday() {
	p.Age++
}

func CallMethod() {
	p := Person{Name: "John", Age: 20}
	p.Birthday()

	fmt.Println(p.Age) // Still 20 - only copy is modified
}

////////////////////////////////////////////////////////////
// Composition
////////////////////////////////////////////////////////////

type Building struct {
	Stories int
	Price   int
}

func (b *Building) PricePerStory() float64 {
	return float64(b.Price) / float64(b.Stories)
}

// Embed the Building type into the House type
type House struct {
	Building
	Detatched bool
}

func Promoted() {
	h := House{
		// Stories: 2 // Doesn't work
		Building: Building{Stories: 2}, // Works
	}

	// Can directly access promoted fields/methods
	h.Stories = 3
	pps := h.PricePerStory()

	fmt.Println(pps)
}

////////////////////////////////////////////////////////////
// Struct Tags
////////////////////////////////////////////////////////////

type User struct {
	ID    int    `json:"user_id"` // Maps {user_id: ""} in JSON to the ID field
	Email string `json:"email" validate:"required"`
}
