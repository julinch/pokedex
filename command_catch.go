package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

const CatchCommandName = "catch"
const CatchCommandDescription = "Cath the pokemon"

const playerXP = 50 //remove?

func GetCatchCommandModel(config *config) cliCommand {
	var command cliCommand
	command.name = CatchCommandName
	command.description = CatchCommandDescription
	command.callback = commandCatch
	return command
}

func commandCatch(config *config, params string) (err error) {
	if len(params) == 0 {
		return errors.New("No pokemon name provided to catch!\n")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params)

	pokemon, err := config.PokeAPIClient.GetPokemon(params, &config.PokeAPIClient.Cache)

	if err != nil {
		return err
	}

	caught := tryCatch(pokemon.BaseExperience)

	if caught {
		fmt.Printf("Success! You caught %s\n", pokemon.PokemonName)
		config.Pokedex.AddCaughtPokemon(pokemon)
		fmt.Printf("Your current caught pokemons are %s !\n", config.Pokedex.GetAllCaughtPokemonNames())
	} else {
		fmt.Printf("Fail! You didn't catch %s\n", pokemon.PokemonName)
	}

	return nil
}

func tryCatch(value int) (isCaught bool) {
	chance := (float64(playerXP) / float64(value))
	rand := rand.Float64()
	// fmt.Printf("chance %v, rand %v", chance, rand)

	return chance > rand
}
