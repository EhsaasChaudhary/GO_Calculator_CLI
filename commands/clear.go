package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearCommand struct
type ClearCommand struct{}

// Execute clears the terminal screen. It does not accept any arguments.
func (c *ClearCommand) Execute(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("❌ The 'clear' command does not accept arguments.\n➡️  Just type 'clear' to clear the screen.")
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

// Description returns a one-line summary of the command.
func (c *ClearCommand) Description() string {
	return "Clears the terminal screen."
}

// Usage returns the usage string for the clear command.
func (c *ClearCommand) Usage() string {
	return "Usage:\n  clear\n\nClears the CLI screen."
}
