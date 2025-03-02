package main

import (
	"GOdemo/commands"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define available commands
	commandMap := map[string]commands.Command{
		"add":   &commands.AddCommand{},
		"sub":   &commands.SubCommand{},
		"mul":   &commands.MulCommand{},
		"div":   &commands.DivCommand{},
		"clear": &commands.ClearCommand{},
		"exit":  &commands.ExitCommand{},
	}

	// Add a help command dynamically
	commandMap["help"] = &commands.HelpCommand{Commands: commandMap}

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("           Welcome to CLI Calculator!         ")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("ℹ️  Type 'help' to see available commands.")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("🔹 Enter command: ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		cmdName := args[0]
		if cmd, exists := commandMap[cmdName]; exists {
			err := cmd.Execute(args[1:])
			if err != nil {
				fmt.Printf("\n🚨 Error: %s\n", err)
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
			}
		} else {
			fmt.Printf("\n❌ Unknown command: '%s'\n➡️  Type 'help' to see available commands.\n", cmdName)
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		}
	}
}
