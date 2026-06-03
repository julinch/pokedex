package pokeapi

type Pokemon struct {
	ID             int    `json:"id"`
	BaseExperience int    `json:"base_experience"`
	PokemonName    string `json:"name"`
}
