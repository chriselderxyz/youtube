package main

import (
	"fmt"
	"maps"
)

////////////////////////////////////////////////////////////
// Working With Maps
////////////////////////////////////////////////////////////

func CreateMaps() {

	// With var syntax
	var m1 map[string]int = map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// With short hand syntax
	m2 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Using make
	// Lets you give a type hint to improve allocation performance
	m3 := make(map[string]int, 10)

	// Zero Value = nil
	// Writing to a nil map = panic
	var m4 map[string]int

	// Empty Map
	// Writing to an empty map = no panic
	// Best practice to use an empty map instead of nil
	var m2 = map[string]int{}
}

func AccessAndModifyMaps() {
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Accessing a key
	v1 := m1["apple"]

	// Modifying a key
	m1["apple"] = 0

	// Adding a new key
	m1["orange"] = 100

	// Checking if a key exists
	a := m1["apple"]         // 0
	k := m1["kiwi"]          // 0
	val1, ok1 := m1["apple"] // ok1 = true (key exists)
	val2, ok2 := m1["kiwi"]  // ok2 = false (key doesn't exist)
}

func DeletingKeyValuePairs() {
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Delete a single key
	delete(m1, "apple")

	// Delete all keys
	clear(m1)
}

func GetNumbKeys() {
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	l := len(m1)
}

func IterateMaps() {
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Iteration happens in a RANDOM ORDER
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}

////////////////////////////////////////////////////////////
// Standard Library maps Package
////////////////////////////////////////////////////////////

func MakeCopy() {
	src := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Make a shallow copy of the map
	m := maps.Clone(src)

	// Merge source into destination
	// Values in source override values in destination
	dst := map[string]int{
		"apple": 10,
	}
	maps.Copy(dst, src)
}

func FilteringMaps() {
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Delete keys that return true
	maps.DeleteFunc(m1, func(key string, value int) bool {
		return value == 2
	})
}

func CompareMaps() {
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	m2 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Check if maps are equal
	eq1 := maps.Equal(m1, m2)

	// Check if maps are equal with a custom logic
	eq2 := maps.EqualFunc(m1, m2, func(v1, v2 int) bool {
		return v1 == v2
	})
}

func IteratorsAndSequences() {
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Get an iterator that iterates over all keys and values
	iter1 := maps.All(m1)

	// Create a map with an iterator
	m2 := maps.Collect(iter1)

	// Insert values from an iterator into a map
	maps.Insert(m1, iter1)

	// Get an iterator that iterates over all keys
	keysIter := maps.Keys(m1)

	// Get an iterator that iterates over all values
	valuesIter := maps.Values(m1)
}
