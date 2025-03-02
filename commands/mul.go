package commands

import (
	"fmt"
	"strconv"
)

// MulCommand struct
type MulCommand struct{}

// Execute multiplies two or more numbers together and prints the result.
// It requires at least two numbers and returns an error if the
// input contains invalid numbers. It prints the result and a
// separator line after the result, and returns an error if
// something went wrong.
func (m *MulCommand) Execute(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("❌ 'mul' requires at least two numbers.\n➡️  Example: mul 2 5 3")
	}

	result := 1.0
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return fmt.Errorf("❌ Invalid number: %s", arg)
		}
		result *= num
	}

	fmt.Printf("\n✅ Result: %v\n", result)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	return nil
}

// Description returns a one-line summary of the command.
func (m *MulCommand) Description() string {
	return "Multiplies multiple numbers."
}

// Usage returns the usage string for the mul command.
func (m *MulCommand) Usage() string {
	return "Usage:\n  mul <num1> <num2> [...]\n\nExamples:\n  mul 2 3 4  - Returns 24"
}
