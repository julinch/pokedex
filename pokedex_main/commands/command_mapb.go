package commands

import (
	"fmt"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

const MapbCommandName = "mapb"
const MapbCommandDescription = "Show previous locations"

func GetMapbCommandModel(Config *Config) CliCommand {
	var command CliCommand
	command.Name = MapbCommandName
	command.Description = MapbCommandDescription
	command.Callback = commandMapb
	return command
}

func commandMapb(Config *Config, params string) (err error) {
	page := pokeapi.Page{}
	if Config == nil || len(Config.Next) == 0 {
		page, err = updateConfig("", Config)

		if err != nil {
			fmt.Print("Issues with poke api %w", err)
			return err
		}
	}

	if Config.Previous == nil {
		fmt.Print("you're on the first page\n")
		return nil
	}

	page, err = updateConfig(*Config.Previous, Config)

	if err != nil {
		return err
	}

	for i := range page.Results {
		fmt.Print(page.Results[i].Name + "\n")
	}
	return nil
}
