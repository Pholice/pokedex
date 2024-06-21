package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Pholice/pokedex/internal/pokecache"
)

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("PokeDex > ")

	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		command, exists := getCommands()[text]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		fmt.Print("\nPokeDex > ")
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	cache pokecache.Cache
	page  int
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Display 20 locations",
			callback:    getMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations",
			callback:    getMapB,
		},
	}
}
