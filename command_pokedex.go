package main

import "fmt"

func pokedex(cfg *config) error {
	fmt.Printf("Your PokedDex:\n")
	for caught := range cfg.pokemon {
		fmt.Printf(" - %s\n", caught)
	}
	return nil
}
