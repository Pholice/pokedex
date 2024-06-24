package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func catch(cfg *config, name string) error {
	var p pokemonPage
	api := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	cached, ok := cfg.cache.Get(api)
	if !ok {
		cached = http_page(cfg, api)
	}
	json.Unmarshal(cached, &p)

	chance := rand.Intn(650) - p.BaseExperience
	fmt.Printf("Base Experience: %d\nThrowing a pokeball at %s...\nRolled: %d\n", p.BaseExperience, name, chance)
	if chance >= p.BaseExperience {
		cfg.pokemon[name] = p
		fmt.Printf("%s was caught!", name)
	} else {
		fmt.Printf("%s got away!", name)
	}
	return nil
}
