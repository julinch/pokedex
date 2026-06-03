package commands

import (
	"fmt"
	"os"
)

const ExitCommandName = "exit"
const ExitCommandDescription = "Exit the Pokedex"

func GetExitCommandModel(Config *Config) CliCommand {
	var command CliCommand
	command.Name = ExitCommandName
	command.Description = ExitCommandDescription
	command.Callback = commandExit
	return command
}

func commandExit(Config *Config, params string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
