package main

import (
	"fmt"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

const MapCommandName = "map"
const MapCommandDescription = "Show locations"

func GetMapCommandModel(config *config) cliCommand {
	var command cliCommand
	command.name = MapCommandName
	command.description = MapCommandDescription
	command.callback = commandMap
	return command
}

func commandMap(config *config, params string) (err error) {
	page := pokeapi.Page{}
	if config == nil || len(config.Next) == 0 {
		page, err = updateConfig("", config)

		if err != nil {
			fmt.Print("Issues with poke api %w", err)
			return err
		}
	} else {
		page, err = updateConfig(config.Next, config)

		if err != nil {
			return err
		}
	}

	for i := range page.Results {
		fmt.Print(page.Results[i].Name + "\n")
	}
	return nil
}
