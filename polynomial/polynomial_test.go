package polynomial

import (
	"testing"
	"math"
)

func TestPolynomialExtrema(t *testing.T) {
	p := Polynomial{
		2: 1,
	}

	zeroes, err := p.Zeroes(0.0001)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !(len(zeroes)==1 && math.Abs(zeroes[0])<0.0001) {
		t.Errorf("Expected [0], got %v", zeroes)
	}

	p[2]=0
	p[4]=1

	zeroes, err = p.Zeroes(0.0001)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !(len(zeroes)==2 && ((math.Abs(zeroes[0]) < 0.0001 && math.Abs(zeroes[1] - 1) < 0.0001) ||
		(math.Abs(zeroes[0] - 1) < 0.0001 && math.Abs(zeroes[1]) < 0.0001))) {
		t.Errorf("Expected [0\t1], got %v", zeroes)
	}
}
