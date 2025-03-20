package main

import (
	"fmt"
	"os"

	"github.com/shais0f/jsonHelper/internal/command"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [arguments]")
		return
	}

	com := os.Args[1]
	args := os.Args[2:]

	if cmd, exists := command.Registry[com]; exists {
		cmd.Execute(args, command.HelpRegistry)
	} else {
		fmt.Println("Unknown command:", com)
		fmt.Println("Use 'help' to see available commands.")
	}
}
