package commands

import (
	"fmt"
	"strconv"
)

// SubCommand struct
type SubCommand struct{}

// Execute subtracts two numbers and prints the result.
// It requires exactly two numbers and returns an error if the
// input contains invalid numbers. It prints the result and a
// separator line after the result, and returns an error if
// something went wrong.
func (s *SubCommand) Execute(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("❌ 'sub' requires exactly two numbers.\n➡️  Example: sub 10 5")
	}

	num1, err1 := strconv.ParseFloat(args[0], 64)
	num2, err2 := strconv.ParseFloat(args[1], 64)
	if err1 != nil || err2 != nil {
		return fmt.Errorf("❌ Invalid numbers")
	}

	result := num1 - num2
	fmt.Printf("\n✅ Result: %v\n", result)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	return nil
}

// Description returns a one-line summary of the sub command.
func (s *SubCommand) Description() string {
	return "Subtracts two numbers."
}

// Usage returns the usage string for the sub command.
func (s *SubCommand) Usage() string {
	return "Usage:\n  sub <num1> <num2>\n\nExamples:\n  sub 10 5   - Returns 5\n  sub 1 2    - Returns -1"
}
