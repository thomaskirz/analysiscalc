package polynomial

import (
	"fmt"
	"strings"
)

type Functions map[string]Polynomial

// Evaluate either stores a new function or returns the string of a requested function
// TODO return error
func (functions Functions) Evaluate(expression string) string {
	if stmt, err := NewParser(strings.NewReader(expression)).Parse(); err == nil {
		if stmt.Function != nil {
			functions[stmt.Name] = stmt.Function
			return fmt.Sprintf("Stored %v(x).", stmt.Name)
		} else {
			if functions[stmt.Name] != nil {
				return functions[stmt.Name].String()
			} else {
				return fmt.Sprintf("Function %v does not exist", stmt.Name)
			}
		}
	} else {
		return err.Error()
	}
}
