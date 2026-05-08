package main

import (
	"bufio"
	"fmt"
	"os"
)

const colorRed = "\033[0;31m"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var config config

	commands := map[string]cliCommand{
		ExitCommandName: GetExitCommandModel(&config),
		HelpCommandName: GetHelpCommandModel(&config),
		MapCommandName:  GetMapCommandModel(&config),
		MapbCommandName: GetMapbCommandModel(&config),
	}

	for _, cmd := range commands {
		UpdateHelpCommandDescription(cmd.name + ": " + cmd.description)
	}

	for {
		fmt.Fprintf(os.Stdout, "%sPokedex > ", colorRed)
		scanner.Scan()
		input := scanner.Text()
		cleanInput := cleanInput(input)

		if len(cleanInput) == 0 {
			fmt.Print("Incorrect input! \n")
		} else {
			if command, ok := commands[cleanInput[0]]; ok {
				command.callback(&config)
			} else {
				fmt.Print("Unknown command \n")
			}
		}
	}
}
