package main

import "fmt"

func Traverse (n *int) {
    // normal loops
    p := *n
	for i:= 0; i <= p; i++ {
		fmt.Println(i)
	}

	

	for p < 35 {
		p++
		n = &p
	}

	fmt.Println(n)
	
}