package pokeapi

type Pokedex struct {
	Pokemons map[string]Pokemon
}

func (pokedex *Pokedex) AddCaughtPokemon(pokemon Pokemon) {
	if pokedex.Pokemons == nil {
		pokedex.Pokemons = make(map[string]Pokemon)
	}

	pokedex.Pokemons[pokemon.Name] = pokemon
}

func (pokedex *Pokedex) GetAllCaughtPokemonNames() (names []string) {

	names = make([]string, 0, len(pokedex.Pokemons))

	for _, p := range pokedex.Pokemons {
		names = append(names, p.Name)
	}

	return names
}
