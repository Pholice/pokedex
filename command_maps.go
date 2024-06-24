package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func getPage(cfg *config) (locationPage, error) {
	var p locationPage
	api := ""
	if cfg.page == 0 {
		api = "https://pokeapi.co/api/v2/location-area/"
	} else {
		api = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", cfg.page*20)
	}

	cached, ok := cfg.cache.Get(api)
	if !ok {
		cached = http_page(cfg, api)
	}
	json.Unmarshal(cached, &p)
	return p, nil
}

func getMap(cfg *config) error {
	cfg.page += 1
	p, err := getPage(cfg)
	if err != nil {
		fmt.Printf("Couldn't get page")
	}
	for _, result := range p.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}

func getMapB(cfg *config) error {
	if cfg.page == 0 {
		return errors.New("nothing behind the first page")
	}
	cfg.page--
	p, err := getPage(cfg)
	if err != nil {
		fmt.Printf("Couldn't get page")
	}
	for _, result := range p.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}
