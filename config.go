package main

import (
	pokeapi "pokedex/internal/poke_api"
	pokecahe "pokedex/internal/pokecache"
	"time"
)

type config struct {
	Previous *string
	Next     string
	Cache    *pokecahe.Cache
}

func updateConfig(url string, config *config) (page pokeapi.Page, err error) {
	if config.Cache == nil {
		config.Cache = pokecahe.NewCache(time.Second * 5)
	}

	page, err = pokeapi.GetPage(url, config.Cache)

	if err != nil {
		return pokeapi.Page{}, err
	}

	config.Previous = page.Previous
	config.Next = page.Next

	return page, nil
}
