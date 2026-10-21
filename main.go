package main

import "fmt"

func main () {
	fmt.Println("Hello From Go")
	n:= 23
	Traverse(&n)
	fmt.Println(n)
	Pointers()
	Slices()
	Interfaces()
}