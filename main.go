package main

import (
	"bufio"
	"fmt"
	"github.com/tombom4/analysiscalc/polynomial"
	"os"
	"strings"
)

var functions = make(map[string]polynomial.Polynomial)

const (
	errunexpected = "An unexpected error occured."
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("analysiscalc command line interface")
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		fmt.Println(evaluate(text))
	}
}

func evaluate(input string) string {
	if stmt, err := polynomial.NewParser(strings.NewReader(input)).Parse(); err != nil {
		return err.Error()
	} else {
		switch stmt.Request {
		case polynomial.ActStore:
			if stmt.Name != "" && stmt.Function != nil {
				functions[stmt.Name] = stmt.Function
				return fmt.Sprintf("Stored function %v(x)", stmt.Name)
			} else {
				return "Empty name or function."
			}
		case polynomial.ActLoad:
			if stmt.Name != "" {
				if functions[stmt.Name] != nil {
					return fmt.Sprintf("%v(x) = %v", stmt.Name, functions[stmt.Name].String())
				} else {
					return fmt.Sprintf("%v(x) does not exist", stmt.Name)
				}
			} else {
				return "Empty name."
			}
		case polynomial.ActDerive:
			if stmt.Name != "" {
				if functions[stmt.Name] != nil {
					return fmt.Sprintf("%v'(x) = %v", stmt.Name, functions[stmt.Name].Derive().String())
				} else {
					return fmt.Sprintf("%v(x) does not exist", stmt.Name)
				}
			} else {
				return "Empty name."
			}
		case polynomial.ActZeroes:
			if stmt.Name != "" {
				if functions[stmt.Name] != nil {
					if zeroes, err := functions[stmt.Name].Zeroes(0.0001); err != nil {
						return err.Error()
					} else {
						var builder strings.Builder
						for _, zero := range zeroes {
							builder.WriteString(fmt.Sprintf("%.3f\t", zero))
						}
						return builder.String()
					}
				} else {
					return fmt.Sprintf("%v(x) does not exist", stmt.Name)
				}
			} else {
				return "Empty name."
			}
		}
		return fmt.Sprint(errunexpected, " ", stmt)
	}
}
