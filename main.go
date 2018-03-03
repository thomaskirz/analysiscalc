package main

import (
	"fmt"
)

func main() {
	// make example Polynomial (-3x^6 + 5x^5 + 33x^3 - 5x^2 + 2x - 1)
	p := Polynomial{
		0: 1,
		1: 3,
		2: -5,
		3: 33,
		5: 5,
		6: -3,
	}

	p.Print()
	fmt.Println(p.Degree())

	// print derivative of p
	fmt.Println("Derivative: ")
	p.Derive().Print()

	// print zeroes of p
	fmt.Printf("Zeroes: %v\n", p.Zeroes())
}
