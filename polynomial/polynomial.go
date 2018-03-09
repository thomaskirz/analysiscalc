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
func (p Polynomial) Zeroes(accuracy float64) ([]float64, error) {

	// TODO: handle saddle points

	if p.Degree() == 1 {
		return []float64{-(p[1] / p[0])}, nil // ax+b=0 => x=-a/b
	} else if p.Degree() == 2 {

		// solve equation with quadratic formula
		discriminant := p[1]*p[1] - (4 * p[2] * p[0]) // determines amount of zeroes
		if discriminant == 0 {
			// the function has one zero
			return []float64{-p[1] / 2 * p[2]}, nil
		} else if discriminant > 0 {
			x1 := (-p[1] - math.Sqrt(discriminant)) / (2 * p[2])
			x2 := (-p[1] + math.Sqrt(discriminant)) / (2 * p[2])
			return []float64{x1, x2}, nil
		} else {
			return []float64{}, nil
		}

	} else {
		extrema, err := p.Extrema(accuracy)
		if err != nil {
			return nil, err
		}

		zeroes := make([]float64, 0)

		// presign and postsign signify if the limit of p is plus or minus infinity (y) at plus and minus infinity (x)
		presign := (p.Degree()%2 == 0 && p[p.Degree()] > 0) || (p.Degree()%2 == 1 && p[p.Degree()] < 0)
		postsign := (p.Degree()%2 == 0 && p[p.Degree()] > 0) || (p.Degree()%2 == 1 && p[p.Degree()] > 0)

		fmt.Printf("extrema: %v\n", extrema)

		for i, extremum := range extrema {
			// if this is the first extremum and there is a zero before it, calculate it with Newton's method
			// if this is the last extremum and there is a zero after it, calculate it with Newton's method
			// otherwise if there is a zero between this and the next extremum, calculate it with Newton's method
			if i == 0 {
				if !math.Signbit(p.Valueat(extremum)) != presign {
					if nwtn, err := p.newton(extremum-1, accuracy); err != nil {
						return nil, err
					} else {
						zeroes = append(zeroes, nwtn)
						fmt.Printf("pre    : i: %v, extremum: %v, zero: %v\n", i, extremum, nwtn)
					}
				}
			}
			if i == len(extrema)-1 {
				if !math.Signbit(p.Valueat(extremum)) != postsign {
					if nwtn, err := p.newton(extremum+1, accuracy); err != nil {
						return nil, err
					} else {
						zeroes = append(zeroes, nwtn)
						fmt.Printf("post   : i: %v, extremum: %v, zero: %v\n", i, extremum, nwtn)
					}
				}
				continue
			}
			if math.Signbit(p.Valueat(extremum)) != math.Signbit(p.Valueat(extrema[i+1])) {
				if nwtn, err := p.newton((extremum+extrema[i+1])/2, accuracy); err != nil {
					return nil, err
				} else {
					zeroes = append(zeroes, nwtn)
					fmt.Printf("between: i: %v, extremum: %v, zero: %v\n", i, extremum, nwtn)
				}
			}
		}

		return zeroes, nil
	}

}

func (p Polynomial) Extrema(accuracy float64) ([]float64, error) {
	return p.Derive().Zeroes(accuracy)
}

func (p Polynomial) Valueat(x float64) (y float64) {
	for exp, coeff := range p {
		y += coeff * math.Pow(x, float64(exp))
	}
	return
}

func (p Polynomial) newton(x float64, e float64) (float64, error) {
	n, m := 0., x
	for i := 0; i < 1000; i++ {
		n = m
		m = n - (p.Valueat(n) / p.Derive().Valueat(n))
		if math.Abs(m-n) < e {
			return m, nil
		}
	}
	return 0, fmt.Errorf("result not converging, starting at %v", x)
}
