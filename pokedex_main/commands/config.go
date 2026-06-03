package commands

import (
	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

type Config struct {
	PokeAPIClient pokeapi.Client
	Pokedex       pokeapi.Pokedex
	Previous      *string
	Next          string
}

func updateConfig(url string, Config *Config) (page pokeapi.Page, err error) {

	page, err = pokeapi.GetPage(url, &Config.PokeAPIClient.Cache)

	if err != nil {
		return pokeapi.Page{}, err
	}

	Config.Previous = page.Previous
	Config.Next = page.Next

	return page, nil
}
