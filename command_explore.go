package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type explorePage struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func exploreResultsPage(cfg *config, location string) (explorePage, error) {
	api := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	var p explorePage
	cached, ok := cfg.cache.Get(api)
	if !ok {
		resp, err := http.Get(api)
		if err != nil {
			fmt.Println("Could not retrieve data from API")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Could not read response body")
		}
		cfg.cache.Add(api, body)
		defer resp.Body.Close()

		err = json.Unmarshal(body, &p)
		if err != nil {
			fmt.Println("Could not unmarshal response body")
		}
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
