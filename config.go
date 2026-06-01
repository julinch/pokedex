package main

import (
	// "time"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
	// pokecahe "github.com/julinch/pokedex/internal/pokecache"
)

type config struct {
	PokeAPIClient pokeapi.Client
	Previous      *string
	Next          string
}

func updateConfig(url string, config *config) (page pokeapi.Page, err error) {
	// if config.PokeAPIClient.Cache == nil {
	// 	config.PokeAPIClient.Cache = pokecahe.NewCache(time.Second * 5)
	// }

	page, err = pokeapi.GetPage(url, &config.PokeAPIClient.Cache)

	if err != nil {
		return pokeapi.Page{}, err
	}

	config.Previous = page.Previous
	config.Next = page.Next

	return page, nil
}
