package main

import (
	"fmt"
	"math"
)

func main() {
	// polynomial represented as slice with it's degree n being the length of the slice - 1
	// the first value is the coefficient of x^n, then of x^(n-1)..., the last value is a constant
	p := []float64{54, 2, -441}
	printPolynomial(p)

	d := derive(p)
	fmt.Println("Derivative:")
	printPolynomial(d)

	if zeroes(p) != nil {
		fmt.Println("\nZero(es):")
		for _, v := range zeroes(p) {
			fmt.Println(v)
		}
	}
}

func printPolynomial(p []float64) {
	degree := len(p) - 1
	first := true
	for _, v := range p {
		if v > 0 {
			// positive coefficient
			if !first {
				fmt.Printf(" + ")
			}
			if degree > 1 {
				fmt.Printf("%gx^%d", v, degree)
			} else if degree == 1 {
				fmt.Printf("%gx", v)
			} else {
				fmt.Printf("%g", v)
			}
			first = false
		} else if v < 0 {
			// negative coefficient
			if !first {
				fmt.Printf(" - ")
			}
			if degree > 1 {
				fmt.Printf("%gx^%d", v*-1, degree)
			} else if degree == 1 {
				fmt.Printf("%gx", v*-1)
			} else {
				fmt.Printf("%g", v*-1)
			}
			first = false
		}
		degree--
	}
	fmt.Println()
	return
}

func derive(function []float64) []float64 {
	degree := len(function) - 1 // degree of function (antiderivative)
	derivative := make([]float64, degree)

	for i := range derivative {
		derivative[i] = function[i] * float64(degree)
		degree--
	}

	return derivative
}

func zeroes(f []float64) []float64 {
	// only works for functions with degree <= 2 for now

	degree := len(f) - 1

	if degree == 1 {
		return []float64{-(f[1] / f[0])} // ax + b = 0 => x = -a/b
	} else if degree == 2 {

		// solve equation with quadratic formula
		discriminant := f[1]*f[1] - (4 * f[0] * f[2]) // determines the amount of zeroes
		if discriminant == 0 {
			// the function has one zero
			return []float64{-f[1] + 2*f[0]}
		} else if discriminant > 0 {
			x1 := (-f[1] - math.Sqrt(discriminant)) / (2 * f[0])
			x2 := (-f[1] + math.Sqrt(discriminant)) / (2 * f[0])
			return []float64{x1, x2}
		}

	}
	return nil
}
