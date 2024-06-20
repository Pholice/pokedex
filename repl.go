package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit PokeDex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("PokeDex > ")

	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		if text == "help" {
			commands[text].callback()
			fmt.Printf("help: %v\n", commands[text].description)
			fmt.Printf("exit: %v\n\n", commands["exit"].description)
		} else if text == "exit" {
			commands[text].callback()
		} else {
			fmt.Printf("Unknown command\n\n")
		}

		fmt.Print("PokeDex > ")
	}
}
