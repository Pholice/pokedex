package main

import "fmt"

func inspect(cfg *config, name string) error {
	stats, exists := cfg.pokemon[name]
	if !exists {
		fmt.Printf("You have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", stats.Height)
	fmt.Printf("Weight: %d\n", stats.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range stats.Stats {
		fmt.Printf("   -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range stats.Types {
		fmt.Printf("   - %s\n", t.Type.Name)
	}
	return nil
}
