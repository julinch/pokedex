package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const colorRed = "\033[0;31m"

func runPokedex(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		ExitCommandName:    GetExitCommandModel(config),
		HelpCommandName:    GetHelpCommandModel(config),
		MapCommandName:     GetMapCommandModel(config),
		MapbCommandName:    GetMapbCommandModel(config),
		ExploreCommandName: GetExploreCommandModel(config),
	}

	for _, cmd := range commands {
		UpdateHelpCommandDescription(cmd.name + ": " + cmd.description)
	}

	for {
		fmt.Fprintf(os.Stdout, "%sPokedex > ", colorRed)
		scanner.Scan()
		input := scanner.Text()
		cleanInput := cleanInput(input)

		switch len(cleanInput) {
		case 0:
			fmt.Print("Incorrect input! \n")
		case 1:
			if command, ok := commands[cleanInput[0]]; ok {
				err := command.callback(config, "")
				if err != nil {
					fmt.Print(err)
				}
			} else {
				fmt.Print("Unknown command \n")
			}

		case 2:
			if command, ok := commands[cleanInput[0]]; ok {
				err := command.callback(config, cleanInput[1])
				if err != nil {
					fmt.Print(err)
				}
			} else {
				fmt.Print("Unknown command \n")
			}

		default:
		}
	}
}

func cleanInput(text string) []string {
	var slice []string
	text = strings.ToLower(text)
	slice = strings.Fields(text)
	return slice
}
