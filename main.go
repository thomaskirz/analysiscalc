package main

import (
	"fmt"
)

func main() {
	// polynomial represented as slice with it's degree n being the length of the slice - 1
	// the first value is the coefficient of x^n, then of x^(n-1)..., the last value is a constant
	p := []uint{3, 5, 0, 3, 5, 2, 1}
	printPolynomial(p)

	d := derive(p)
	fmt.Println("Derivative:")
	printPolynomial(d)
}

func printPolynomial(p []uint) {
	degree := len(p) - 1
	for _, v := range p {
		if v != 0 {
			if degree == len(p)-1 {
				fmt.Printf("%dx^%d", v, degree)
			} else if degree > 1 {
				fmt.Printf(" + %dx^%d", v, degree)
			} else if degree == 1 {
				fmt.Printf(" + %dx", v)
			} else {
				fmt.Printf(" + %d", v)
			}
		}
		degree--
	}
	fmt.Println()
	return
}

func derive(function []uint) []uint {
	degree := uint(len(function) - 1) // degree of derivative
	derivative := make([]uint, degree)

	for i := range derivative {
		derivative[i] = function[i] * degree
		degree--
	}

	return derivative
}
