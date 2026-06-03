package commands

import (
	"fmt"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

const MapCommandName = "map"
const MapCommandDescription = "Show locations"

func GetMapCommandModel(Config *Config) CliCommand {
	var command CliCommand
	command.Name = MapCommandName
	command.Description = MapCommandDescription
	command.Callback = commandMap
	return command
}

func commandMap(Config *Config, params string) (err error) {
	page := pokeapi.Page{}
	if Config == nil || len(Config.Next) == 0 {
		page, err = updateConfig("", Config)

		if err != nil {
			fmt.Print("Issues with poke api %w", err)
			return err
		}
	} else {
		page, err = updateConfig(Config.Next, Config)

		if err != nil {
			return err
		}
	}

	for i := range page.Results {
		fmt.Print(page.Results[i].Name + "\n")
	}
	return nil
}
