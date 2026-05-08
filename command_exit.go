package main

import (
	"fmt"
	"os"
)

const ExitCommandName = "exit"
const ExitCommandDescription = "Exit the Pokedex"

func GetExitCommandModel(config *config) cliCommand {
	var command cliCommand
	command.name = ExitCommandName
	command.description = ExitCommandDescription
	command.callback = commandExit
	return command
}

func commandExit(config *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
