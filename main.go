package main

import (
	"time"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
	pokedexmain "github.com/julinch/pokedex/pokedex_main"
	"github.com/julinch/pokedex/pokedex_main/commands"
)

func main() {

	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &commands.Config{
		PokeAPIClient: client,
	}
	pokedexmain.RunPokedex(config)
}
