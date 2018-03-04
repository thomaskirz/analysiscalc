package main

import (
	"bufio"
	"fmt"
	"github.com/tombom4/analysiscalc/polynomial"
	"os"
)

func main() {
	functions := make(polynomial.Functions)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Waiting for function: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(functions.Evaluate(text))
	}
}
