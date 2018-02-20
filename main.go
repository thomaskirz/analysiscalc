package main

import (
	"fmt"
)

// Polynomial represents a polynomial function
type Polynomial struct {
	Degree   uint
	Summands map[uint]float64
}

func main() {
	p := Polynomial{
		6,
		map[uint]float64{
			0: -1,
			1: 2,
			2: -5,
			3: 33,
			5: 5,
			6: -3,
		},
	}

	printPolynomial(p)
}

func printPolynomial(p Polynomial) {
	s := p.Summands
	deg := p.Degree

	// iterate down from highest (degree) to lowest (0) coefficient
	for i := deg; ; i-- {
		if s[i] > 0 {
			// positive coefficient
			if i < deg {
				fmt.Printf(" + ")
			}
			if i > 1 {
				fmt.Printf("%gx^%d", s[i], i)
			} else if i == 1 {
				fmt.Printf("%gx", s[i])
			} else {
				fmt.Printf("%g", s[i])
			}
		} else if s[i] < 0 {
			// negative coefficient
			if i < deg {
				fmt.Printf(" - ")
			} else {
				fmt.Printf("-")
			}
			if i > 1 {
				fmt.Printf("%gx^%d", s[i]*-1, i)
			} else if i == 1 {
				fmt.Printf("%gx", s[i]*-1)
			} else {
				fmt.Printf("%g", s[i]*-1)
			}
		}
		if i == 0 {
			break
		}
	}
	fmt.Println()
}

func derive(function []uint) []uint {
	// TODO

	// degree := uint(len(function) - 1) // degree of function
	// derivative := make([]uint, degree)

	// for i := range derivative {
	// 	derivative[i] = function[i] * degree
	// 	degree--
	// }

	// return derivative
}
