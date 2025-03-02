package commands

import (
	"fmt"
	"os"
)

// ExitCommand struct
type ExitCommand struct{}

// Execute exits the CLI application immediately. It does not accept any arguments.
func (e *ExitCommand) Execute(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("❌ The 'exit' command does not accept arguments.\n➡️  Just type 'exit' to quit.")
	}

	fmt.Println("\n👋 Exiting... Goodbye!\n")
	os.Exit(0)
	return nil
}

// Description returns a one-line summary of the exit command.
func (e *ExitCommand) Description() string {
	return "Exits the application."
}

// Usage returns the usage string for the exit command.
func (e *ExitCommand) Usage() string {
	return "Usage:\n  exit\n\nExits the CLI application immediately."
}
