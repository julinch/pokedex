package main

import (
	"fmt"
	pokeapi "pokedex/internal/poke_api"
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

func commandMap(config *config) (err error) {
	page := pokeapi.Page{}
	if config == nil || len(config.Next) == 0 {
		fmt.Print("MAP config or config next nil\n\n\n")
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
