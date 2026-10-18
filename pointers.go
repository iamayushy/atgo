package main

import "fmt"

type Counter struct {
	count int
}

// pointer receiver: changes are visible to the caller
func (c *Counter) Increment() {
	c.count++
}

// value receiver: works on a copy
func (c Counter) Value() int {
	return c.count
}

// value parameter: the caller's variable is not changed
func doubleValue(x int) {
	x = x * 2
}

// pointer parameter: the caller's variable is changed
func doublePointer(x *int) {
	*x = *x * 2
}

func Pointers() {
	// basics: & takes the address, * reads/writes through it
	a := 10
	p := &a
	fmt.Println("a:", a, "address:", p, "value at p:", *p)

	*p = 20
	fmt.Println("after *p = 20, a:", a)

	// zero value of a pointer is nil
	var np *int
	fmt.Println("nil pointer:", np, np == nil)

	// new allocates a zeroed value and returns a pointer to it
	q := new(int)
	*q = 5
	fmt.Println("new(int):", *q)

	// pass by value vs pass by pointer
	b := 7
	doubleValue(b)
	fmt.Println("after doubleValue:", b)
	doublePointer(&b)
	fmt.Println("after doublePointer:", b)

	// pointers to structs: fields are accessed without explicit *
	c := &Counter{}
	c.Increment()
	c.Increment()
	fmt.Println("counter:", c.Value())

	// pointer to pointer
	pp := &p
	**pp = 30
	fmt.Println("after **pp = 30, a:", a)
}
