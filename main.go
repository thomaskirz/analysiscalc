package main

import (
	"fmt"
	"math"
)

// Polynomial represents a polynomial function.
// Degree must be the highest degree of the summands of the function.
type Polynomial struct {
	Degree   uint
	Summands map[uint]float64
}

func main() {
	// make example Polynomial (-3x^6 + 5x^5 + 33x^3 - 5x^2 + 2x - 1)
	p := Polynomial{
		2,
		map[uint]float64{
			0: 1,
			1: 3,
			2: -5,
			// 3: 33,
			// 5: 5,
			// 6: -3,
		},
	}

	PrintPolynomial(p)

	// print derivative of p
	fmt.Println("Derivative: ")
	PrintPolynomial(Derive(p))

	// print zeroes of p
	fmt.Printf("Zeroes: %v\n", Zeroes(p))
}

// PrintPolynomial prints a polynomial in the form `5x^5 + 18x^3 - 5x^2 + 2x -1`.
func PrintPolynomial(p Polynomial) {
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

// Derive returns the derivative of a Polynomial as Polynomial
func Derive(p Polynomial) Polynomial {
	// this map stores the derivative
	dsumms := make(map[uint]float64, len(p.Summands))

	for i := p.Degree; i > 0; i-- {
		if p.Summands[i] != 0 {
			dsumms[i-1] = p.Summands[i] * float64(i)
		}
	}

	return Polynomial{uint(len(dsumms)), dsumms}
}

// Zeroes returns a slice containing the zeroes of a quadratic of linear function in no specific order.
func Zeroes(p Polynomial) []float64 {

	if p.Degree == 1 {
		return []float64{-(p.Summands[1] / p.Summands[0])} // ax+b=0 => x=-a/b
	} else if p.Degree == 2 {
		// solve equation with quadratic formula
		discriminant := p.Summands[1]*p.Summands[1] - (4 * p.Summands[2] * p.Summands[0]) // determines amount of zeroes
		if discriminant == 0 {
			// the function has one zero
			return []float64{-p.Summands[1] / 2 * p.Summands[2]}
		} else if discriminant > 0 {
			x1 := (-p.Summands[1] - math.Sqrt(discriminant)) / (2 * p.Summands[2])
			x2 := (-p.Summands[1] + math.Sqrt(discriminant)) / (2 * p.Summands[2])
			return []float64{x1, x2}
		}
	}

	return nil
}
