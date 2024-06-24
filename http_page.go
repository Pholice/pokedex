package main

import (
	"fmt"
	"io"
	"net/http"
)

func http_page(cfg *config, api string) []byte {
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

	return body
}
