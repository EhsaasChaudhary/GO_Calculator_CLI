package commands

import (
	"fmt"
)

// Command interface
type Command interface {
	Execute(args []string) error
	Description() string
	Usage() string
}

// HelpCommand provides detailed command help
type HelpCommand struct {
	Commands map[string]Command
}

// Execute provides detailed help for a command, or a list of all commands.
func (h *HelpCommand) Execute(args []string) error {
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("                  HELP MENU                   ")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	if len(args) > 0 {
		cmdName := args[0]
		if cmd, exists := h.Commands[cmdName]; exists {
			fmt.Printf("\n%s\n", cmd.Usage())
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			return nil
		}
		return fmt.Errorf("❌ Unknown command: '%s'\n➡️  Type 'help' to see available commands.", cmdName)
	}

	fmt.Println("Available commands:\n")
	for name, cmd := range h.Commands {
		fmt.Printf("  %-10s - %s\n", name, cmd.Description())
	}
	fmt.Println("\nℹ️  Use 'help <command>' for detailed command info.")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	return nil
}

// Description returns a one-line summary of the command.
func (h *HelpCommand) Description() string {
	return "Displays a list of commands or help for a specific command"
}

// Usage returns the usage string for the help command.
func (h *HelpCommand) Usage() string {
	return "Usage:\n  help [command]\n\nExamples:\n  help           - Lists all available commands\n  help add       - Displays help for 'add' command"
}
