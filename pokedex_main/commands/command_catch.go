package commands

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

const CatchCommandName = "catch"
const CatchCommandDescription = "Cath the pokemon"

const playerXP = 50 //remove?

func GetCatchCommandModel(Config *Config) CliCommand {
	var command CliCommand
	command.Name = CatchCommandName
	command.Description = CatchCommandDescription
	command.Callback = commandCatch
	return command
}

func commandCatch(Config *Config, params string) (err error) {
	if len(params) == 0 {
		return errors.New("No pokemon name provided to catch!\n")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params)

	pokemon, err := Config.PokeAPIClient.GetPokemon(params, &Config.PokeAPIClient.Cache)

	if err != nil {
		return err
	}

	caught := tryCatch(pokemon.BaseExperience)

	if caught {
		fmt.Printf("Success! You caught %s\n", pokemon.Name)
		Config.Pokedex.AddCaughtPokemon(pokemon)
		fmt.Print("You may now inspect it with the inspect command.\n")
	} else {
		fmt.Printf("Fail! You didn't catch %s\n", pokemon.Name)
	}

	return nil
}

func tryCatch(value int) (isCaught bool) {
	chance := (float64(playerXP) / float64(value))
	rand := rand.Float64()
	// fmt.Printf("chance %v, rand %v", chance, rand)

	return chance > rand
}
