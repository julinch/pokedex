package commands

import (
	"fmt"
	"strings"

	pokeapi "github.com/julinch/pokedex/internal/poke_api"
)

const InspectCommandName = "inspect"
const InspectCommandDescription = "Inspect the pokemon's stats"

func GetInspectCommandModel(Config *Config) CliCommand {
	var command CliCommand
	command.Name = InspectCommandName
	command.Description = InspectCommandDescription
	command.Callback = commandInspect
	return command
}

func commandInspect(Config *Config, params string) (err error) {

	if len(params) == 0 {
		return fmt.Errorf("Pokemon name is empty!\n")
	}

	if pokemon, ok := Config.Pokedex.Pokemons[params]; !ok {
		fmt.Printf("You have not caught %s yet\n", params)
	} else {
		fmt.Print(GetPokemonStatsString(pokemon))
	}
	return nil
}

func GetPokemonStatsString(pokemon pokeapi.Pokemon) string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "Name: %s\n", pokemon.Name)
	fmt.Fprintf(&sb, "Height: %d\n", pokemon.Height)
	fmt.Fprintf(&sb, "Weight: %d\n", pokemon.Weight)

	sb.WriteString("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Fprintf(&sb, "  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	sb.WriteString("Types:\n")
	for _, pokemonType := range pokemon.Types {
		fmt.Fprintf(&sb, "  - %s\n", pokemonType.Type.Name)
	}

	return sb.String()
}
