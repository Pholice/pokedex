package main

import (
	"encoding/json"
	"fmt"
)

func exploreResultsPage(cfg *config, location string) (explorePage, error) {
	api := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	var p explorePage
	cached, ok := cfg.cache.Get(api)
	if !ok {
		cached = http_page(cfg, api)
	}
	json.Unmarshal(cached, &p)

	return p, nil
}

func explore(cfg *config, location string) error {
	p, err := exploreResultsPage(cfg, location)
	if err != nil {
		fmt.Printf("Couldn't get page")
	}
	for _, pokemon := range p.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}

	return nil
}
