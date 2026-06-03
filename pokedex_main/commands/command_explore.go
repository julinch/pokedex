package commands

import (
	"errors"
	"fmt"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

const ExploreCommandName = "explore"
const ExploreCommandDescription = "Explore the location"

func GetExploreCommandModel(Config *Config) CliCommand {
	var command CliCommand
	command.Name = ExploreCommandName
	command.Description = ExploreCommandDescription
	command.Callback = commandExplore
	return command
}

func commandExplore(Config *Config, params string) (err error) {
	if len(params) == 0 {
		return errors.New("Empty location name!\n")
	}

	area := pokeapi.LocationArea{}

	if Config == nil {
		return fmt.Errorf("MAP Config nil\n\n\n")
	}

	area, err = Config.PokeAPIClient.GetLocationArea(params, &Config.PokeAPIClient.Cache)

	if err != nil {
		return err
	}

	for i := 0; i < len(area.PokemonEncounters); i++ {
		encounter := area.PokemonEncounters[i]
		fmt.Printf("%s", encounter.Pokemon.Name+"\n")
	}
	fmt.Printf("%s", area.Name)

	return nil
}
