package main

import "fmt"

// slices share their backing array: changes are visible to the caller
func fillZero(s []int) {
	for i := range s {
		s[i] = 0
	}
}

// append may allocate a new array, so return the slice
func addItem(s []int, v int) []int {
	return append(s, v)
}

func Slices() {
	// arrays have a fixed size that is part of the type
	arr := [3]int{1, 2, 3}
	fmt.Println("array:", arr, "len:", len(arr))

	// slices are a view over an array: pointer, length, capacity
	s := []int{10, 20, 30, 40, 50}
	fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))

	// slicing s[low:high] includes low, excludes high
	part := s[1:3]
	fmt.Println("s[1:3]:", part, "len:", len(part), "cap:", cap(part))

	// sub-slices share memory with the original
	part[0] = 99
	fmt.Println("after part[0] = 99, s:", s)

	// make creates a slice with a given length and capacity
	m := make([]int, 0, 2)
	for i := 1; i <= 4; i++ {
		m = addItem(m, i)
		fmt.Println("append", i, "len:", len(m), "cap:", cap(m))
	}

	// copy makes an independent slice
	c := make([]int, len(s))
	copy(c, s)
	c[0] = -1
	fmt.Println("copy:", c, "original:", s)

	// passing a slice lets the function change its elements
	fillZero(part)
	fmt.Println("after fillZero(part), s:", s)

	// zero value of a slice is nil, but len and append still work
	var ns []int
	fmt.Println("nil slice:", ns == nil, "len:", len(ns))
	ns = append(ns, 1)
	fmt.Println("after append:", ns)

	// maps: key/value pairs
	ages := map[string]int{"alice": 30, "bob": 25}
	ages["carol"] = 35
	fmt.Println("map:", ages, "len:", len(ages))

	// missing keys return the zero value; use comma ok to check
	if v, ok := ages["dave"]; !ok {
		fmt.Println("dave missing, zero value:", v)
	}

	delete(ages, "bob")
	fmt.Println("after delete bob:", ages)

	// map iteration order is random
	for name, age := range ages {
		fmt.Println(name, "is", age)
	}
}
