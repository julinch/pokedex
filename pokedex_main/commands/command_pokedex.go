package commands

import (
	"fmt"
)

const PokedexCommandName = "pokedex"
const PokedexCommandDescription = "See the list of all caught pokemons"

func GetPokedexCommandModel(Config *Config) CliCommand {
	var command CliCommand
	command.Name = PokedexCommandName
	command.Description = PokedexCommandDescription
	command.Callback = commandPokedex
	return command
}

func commandPokedex(Config *Config, params string) (err error) {
	pokemonNames := Config.Pokedex.GetAllCaughtPokemonNames()

	if len(pokemonNames) == 0 {
		fmt.Print("Your pokedex is empty!")
		return nil
	}

	fmt.Printf("Your pokedex:\n")

	for _, pokemonName := range pokemonNames {
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}
