package main

import (
	"fmt"
	"github.com/tombom4/analysiscalc/polynomial"
)

func main() {
	//functions := make(polynomial.Functions)
	//
	//reader := bufio.NewReader(os.Stdin)
	//for {
	//	fmt.Print("Waiting for function: ")
	//	text, _ := reader.ReadString('\n')
	//	fmt.Println(functions.Evaluate(text))
	//}

	p := polynomial.Polynomial{
		0: 1,
		1: 1,
		2: -3,
		3: 0,
		4: 1,
	}

	zeroes, _ := p.Zeroes(0.001)
	for _, v := range zeroes {
		fmt.Printf("%.3f\t", v)
	}
	fmt.Println()
}
