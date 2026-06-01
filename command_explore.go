package main

import (
	"errors"
	"fmt"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

const ExploreCommandName = "explore"
const ExploreCommandDescription = "Explore the location"

func GetExploreCommandModel(config *config) cliCommand {
	var command cliCommand
	command.name = ExploreCommandName
	command.description = ExploreCommandDescription
	command.callback = commandExplore
	return command
}

func commandExplore(config *config, params string) (err error) {
	if len(params) == 0 {
		return errors.New("Empty location name!")
	}

	area := pokeapi.LocationArea{}

	if config == nil {
		fmt.Print("MAP config nil\n\n\n")

	}

	area, err = config.PokeAPIClient.GetLocationArea(params, &config.PokeAPIClient.Cache)

	if err != nil {
		return err
	}

	for i := 0; i < len(area.PokemonEncounters); i++ {
		encounter := area.PokemonEncounters[i]
		fmt.Printf(encounter.Pokemon.Name + "\n")
	}
	fmt.Printf(area.Name)

	return nil
}
