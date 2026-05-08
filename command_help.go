package main

import (
	"fmt"
)

const HelpCommandName = "help"
const HelpCommandDescription = "Displays a help message"

var AllCommandsDescriptions = "Welcome to the Pokedex!\nUsage:\n\n"

func GetHelpCommandModel(config *config) cliCommand {
	var command cliCommand
	command.name = HelpCommandName
	command.description = HelpCommandDescription
	command.callback = commandHelp
	return command
}

func UpdateHelpCommandDescription(commandDescription string) {
	AllCommandsDescriptions += commandDescription + "\n"
}

func commandHelp(config *config) error {
	fmt.Print(AllCommandsDescriptions)
	return nil
}
