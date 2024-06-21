package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type page struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getPageAPI(cfg *config) string {
	if cfg.page == 0 {
		return "https://pokeapi.co/api/v2/location-area/"
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", cfg.page*20)
	return url
}

func getPage(cfg *config) (page, error) {
	var p page
	cached, ok := cfg.cache.Get(getPageAPI(cfg))
	if !ok {
		resp, err := http.Get(getPageAPI(cfg))
		if err != nil {
			fmt.Println("Could not retrieve data from API")
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Could not read response body")
		}
		cfg.cache.Add(getPageAPI(cfg), body)
		defer resp.Body.Close()

		err = json.Unmarshal(body, &p)
		if err != nil {
			fmt.Println("Could not unmarshal response body")
		}
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
