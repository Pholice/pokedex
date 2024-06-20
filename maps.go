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

func getPageAPI() string {
	if pageIndex == 0 {
		return "https://pokeapi.co/api/v2/location-area/"
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", pageIndex*20)
	return url
}

func getPage() (page, error) {
	var p page
	resp, err := http.Get(getPageAPI())
	if err != nil {
		fmt.Println("Could not retrieve data from API")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read response body")
	}

	defer resp.Body.Close()
	err = json.Unmarshal(body, &p)
	if err != nil {
		fmt.Println("Could not unmarshal response body")
	}

	return p, nil
}

func getMap() error {
	pageIndex += 1
	p, err := getPage()
	if err != nil {
		fmt.Printf("Couldn't get page")
	}
	for _, result := range p.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}

func getMapB() error {
	if pageIndex == 0 {
		return errors.New("nothing behind the first page")
	}
	pageIndex--
	p, err := getPage()
	if err != nil {
		fmt.Printf("Couldn't get page")
	}
	for _, result := range p.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}
