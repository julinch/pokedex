package main

import (
	"fmt"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

const MapbCommandName = "mapb"
const MapbCommandDescription = "Show previous locations"

func GetMapbCommandModel(config *config) cliCommand {
	var command cliCommand
	command.name = MapbCommandName
	command.description = MapbCommandDescription
	command.callback = commandMapb
	return command
}

func commandMapb(config *config, params string) (err error) {
	page := pokeapi.Page{}
	if config == nil || len(config.Next) == 0 {
		page, err = updateConfig("", config)

		if err != nil {
			fmt.Print("Issues with poke api %w", err)
			return err
		}
	}

	if config.Previous == nil {
		fmt.Print("you're on the first page\n")
		return nil
	}

	page, err = updateConfig(*config.Previous, config)

	if err != nil {
		return err
	}

	for i := range page.Results {
		fmt.Print(page.Results[i].Name + "\n")
	}
	return nil
}
