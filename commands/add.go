package commands

import (
	"fmt"
	"strconv"
)

// AddCommand struct
type AddCommand struct{}

// Execute adds several numbers together and prints the result.
// It requires at least two numbers and returns an error if the
// input contains invalid numbers. It prints the result and a
// separator line after the result, and returns an error if
// something went wrong.
func (a *AddCommand) Execute(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("❌ 'add' requires at least two numbers.\n➡️  Example: add 2 5 8")
	}

	var sum float64
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return fmt.Errorf("❌ Invalid number: %s", arg)
		}
		sum += num
	}

	fmt.Printf("\n✅ Result: %v\n", sum)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	return nil
}

// Description returns a one-line summary of the command.
func (a *AddCommand) Description() string {
	return "Adds multiple numbers."
}

// Usage returns the usage string for the add command.
func (a *AddCommand) Usage() string {
	return "Usage:\n  add <num1> <num2> [...]\n\nExamples:\n  add 2 5 8   - Returns 15\n  add 1.5 2.5 - Returns 4"
}
