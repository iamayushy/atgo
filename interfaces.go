package main

import (
	"fmt"
	"math"
)

// an interface lists method signatures; any type with those methods satisfies it
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// no "implements" keyword: Rect satisfies Shape just by having the methods
func (r Rect) Area() float64      { return r.width * r.height }
func (r Rect) Perimeter() float64 { return 2 * (r.width + r.height) }

func (c Circle) Area() float64      { return math.Pi * c.radius * c.radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.radius }

// fmt.Stringer is a standard interface: String() controls how a value prints
func (c Circle) String() string {
	return fmt.Sprintf("Circle(r=%.1f)", c.radius)
}

// functions can accept any type that satisfies the interface
func describe(s Shape) {
	fmt.Printf("%v area: %.2f perimeter: %.2f\n", s, s.Area(), s.Perimeter())
}

func Interfaces() {
	// a slice of interface values can hold different concrete types
	shapes := []Shape{Rect{3, 4}, Circle{2}}
	for _, s := range shapes {
		describe(s)
	}

	// type assertion: get the concrete type back out
	var s Shape = Circle{5}
	if c, ok := s.(Circle); ok {
		fmt.Println("it's a circle with radius", c.radius)
	}

	// type switch: branch on the concrete type
	for _, s := range shapes {
		switch v := s.(type) {
		case Rect:
			fmt.Println("rect", v.width, "x", v.height)
		case Circle:
			fmt.Println("circle", v.radius)
		}
	}

	// the empty interface (any) accepts values of any type
	values := []any{42, "hello", true, Rect{1, 2}}
	for _, v := range values {
		fmt.Printf("%v is %T\n", v, v)
	}

	// a nil interface has no type and no value
	var empty Shape
	fmt.Println("nil interface:", empty == nil)
}
