package main

import (
	// "time"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

type config struct {
	PokeAPIClient pokeapi.Client
	// Pokemons      map[string]pokeapi.Pokemon
	Pokedex  pokeapi.Pokedex
	Previous *string
	Next     string
}

func updateConfig(url string, config *config) (page pokeapi.Page, err error) {

	page, err = pokeapi.GetPage(url, &config.PokeAPIClient.Cache)

	if err != nil {
		return pokeapi.Page{}, err
	}

	config.Previous = page.Previous
	config.Next = page.Next

	return page, nil
}
