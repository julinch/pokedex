package main

import (
	"time"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

func main() {

	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &config{
		PokeAPIClient: client,
	}
	runPokedex(config)
}
