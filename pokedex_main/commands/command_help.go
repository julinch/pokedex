package commands

import (
	"fmt"
)

const HelpCommandName = "help"
const HelpCommandDescription = "Displays a help message"

var AllCommandsDescriptions = "Welcome to the Pokedex!\nUsage:\n\n"

func GetHelpCommandModel(Config *Config) CliCommand {
	var command CliCommand
	command.Name = HelpCommandName
	command.Description = HelpCommandDescription
	command.Callback = commandHelp
	return command
}

func UpdateHelpCommandDescription(commandDescription string) {
	AllCommandsDescriptions += commandDescription + "\n"
}

func commandHelp(Config *Config, params string) error {
	fmt.Print(AllCommandsDescriptions)
	return nil
}
