package commands

import (
	"fmt"
	"strconv"
)

// DivCommand struct
type DivCommand struct{}

// Execute divides two numbers and prints the result.
// It requires exactly two numbers and returns an error if the
// input contains invalid numbers. It prints the result and a
// separator line after the result, and returns an error if
// something went wrong.
func (d *DivCommand) Execute(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("❌ 'div' requires exactly two numbers.\n➡️  Example: div 10 5")
	}

	num1, err1 := strconv.ParseFloat(args[0], 64)
	num2, err2 := strconv.ParseFloat(args[1], 64)
	if err1 != nil || err2 != nil {
		return fmt.Errorf("❌ Invalid numbers.")
	}

	if num2 == 0 {
		return fmt.Errorf("🚨 Division by zero is not allowed!")
	}

	result := num1 / num2
	fmt.Printf("\n✅ Result: %v\n", result)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	return nil
}

// Description returns a one-line summary of the div command.
func (d *DivCommand) Description() string {
	return "Divides two numbers."
}

// Usage returns the usage string for the div command.
func (d *DivCommand) Usage() string {
	return "Usage:\n  div <num1> <num2>\n\nExamples:\n  div 10 2   - Returns 5\n  div 1 2    - Returns 0.5"
}
