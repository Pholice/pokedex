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
		words := strings.Split(text, " ")
		command, exists := getCommands()[words[0]]
		if exists {
			if command.callbackarg != nil {
				err := command.callbackarg(cfg, words[1])
				if err != nil {
					fmt.Printf("Error executing callbackarg")
				}
			}
			if command.callback != nil {
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
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
	callbackarg func(*config, string) error
}

type config struct {
	cache   pokecache.Cache
	page    int
	pokemon map[string]pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
			callbackarg: nil,
		},
		"exit": {
			name:        "exit",
			description: "Exit PokeDex",
			callback:    commandExit,
			callbackarg: nil,
		},
		"map": {
			name:        "map",
			description: "Display 20 locations",
			callback:    getMap,
			callbackarg: nil,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations",
			callback:    getMapB,
			callbackarg: nil,
		},
		"explore": {
			name:        "explore",
			description: "explore <location_name> - Explores location for Pokemon",
			callback:    nil,
			callbackarg: explore,
		},
		"catch": {
			name:        "catch",
			description: "catch <pokemon_name> - Attempts to catch Pokemon",
			callback:    nil,
			callbackarg: catch,
		},
	}
}
