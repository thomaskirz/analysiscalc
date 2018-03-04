package polynomial

import (
	"fmt"
	"math"
	"strings"
)

// Polynomial represents a polynomial function.
type Polynomial map[uint]float64

// Degree returns the (highest) degree of the polynomial.
func (p Polynomial) Degree() (deg uint) {
	for key, value := range p {
		if value != 0 {
			deg = uint(math.Max(float64(key), float64(deg)))
		}
	}
	return
}

// String returns a polynomial as a string in the form '5x^5 + 18x^3 - 5x^2 + 2x -1'.
func (p Polynomial) String() string {
	var str strings.Builder

	// iterate down from highest (degree) to lowest (0) coefficient
	for i := p.Degree(); ; i-- {
		if p[i] > 0 {
			// positive coefficient
			if i < p.Degree() {
				str.WriteString(" + ")
			}
			if i > 1 {
				str.WriteString(fmt.Sprintf("%gx^%d", p[i], i))
			} else if i == 1 {
				str.WriteString(fmt.Sprintf("%gx", p[i]))
			} else {
				str.WriteString(fmt.Sprintf("%g", p[i]))
			}
		} else if p[i] < 0 {
			// negative coefficient
			if i < p.Degree() {
				str.WriteString(" - ")
			} else {
				str.WriteString("-")
			}
			if i > 1 {
				str.WriteString(fmt.Sprintf("%gx^%d", p[i]*-1, i))
			} else if i == 1 {
				str.WriteString(fmt.Sprintf("%gx", p[i]*-1))
			} else {
				str.WriteString(fmt.Sprintf("%g", p[i]*-1))
			}
		}
		if i == 0 {
			break
		}
	}
	return str.String()
}

// Print prints a polynomial in the form '5x^5 + 18x^3 - 5x^2 + 2x -1'.
func (p Polynomial) Print() {
	fmt.Println(p.String())
}

// Derive returns the derivative of a Polynomial as Polynomial
func (p Polynomial) Derive() Polynomial {
	// this map stores the derivative
	derivative := make(Polynomial)

	for i := p.Degree(); i > 0; i-- {
		if p[i] != 0 {
			derivative[i-1] = p[i] * float64(i)
		}
	}

	return derivative
}

// Zeroes returns a slice containing the zeroes of a quadratic of linear function in no specific order.
func (p Polynomial) Zeroes() []float64 {

	if p.Degree() == 1 {
		return []float64{-(p[1] / p[0])} // ax+b=0 => x=-a/b
	} else if p.Degree() == 2 {
		// solve equation with quadratic formula
		discriminant := p[1]*p[1] - (4 * p[2] * p[0]) // determines amount of zeroes
		if discriminant == 0 {
			// the function has one zero
			return []float64{-p[1] / 2 * p[2]}
		} else if discriminant > 0 {
			x1 := (-p[1] - math.Sqrt(discriminant)) / (2 * p[2])
			x2 := (-p[1] + math.Sqrt(discriminant)) / (2 * p[2])
			return []float64{x1, x2}
		}
	}

	return nil
}
