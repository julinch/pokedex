package pokedexmain

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commands "github.com/julinch/pokedex/pokedex_main/commands"
)

const colorRed = "\033[0;31m"

func RunPokedex(Config *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	commandsList := map[string]commands.CliCommand{
		commands.ExitCommandName:    commands.GetExitCommandModel(Config),
		commands.HelpCommandName:    commands.GetHelpCommandModel(Config),
		commands.MapCommandName:     commands.GetMapCommandModel(Config),
		commands.MapbCommandName:    commands.GetMapbCommandModel(Config),
		commands.ExploreCommandName: commands.GetExploreCommandModel(Config),
		commands.CatchCommandName:   commands.GetCatchCommandModel(Config),
		commands.InspectCommandName: commands.GetInspectCommandModel(Config),
		commands.PokedexCommandName: commands.GetPokedexCommandModel(Config),
	}

	for _, cmd := range commandsList {
		commands.UpdateHelpCommandDescription(cmd.Name + ": " + cmd.Description)
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
			if command, ok := commandsList[cleanInput[0]]; ok {
				err := command.Callback(Config, "")
				if err != nil {
					fmt.Print(err)
				}
			} else {
				fmt.Print("Unknown command \n")
			}

		case 2:
			if command, ok := commandsList[cleanInput[0]]; ok {
				err := command.Callback(Config, cleanInput[1])
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
