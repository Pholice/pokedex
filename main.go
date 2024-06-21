package main

import (
	"time"

	"github.com/Pholice/pokedex/internal/pokecache"
)

func main() {
	clearInterval := 5 * time.Minute
	cfg := &config{
		cache:   pokecache.NewCache(clearInterval),
		page:    -1,
		pokemon: make(map[string]pokemon),
	}
	repl(cfg)
}
