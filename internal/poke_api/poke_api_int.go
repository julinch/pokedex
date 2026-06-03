package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	pokecahe "github.com/julinch/pokedex/internal/pokecache"
)

const pokedexLocationAreaEP = "https://pokeapi.co/api/v2/location-area/"
const pokedexPokemonEP = "https://pokeapi.co/api/v2/pokemon/"

func GetPage(areaName string, cache *pokecahe.Cache) (page Page, err error) {

	if entry, is := cache.Get(areaName); is {
		err = json.Unmarshal(entry, &page)

		if err != nil {
			return Page{}, err
		}

		return page, nil
	}

	if len(areaName) == 0 {
		areaName = pokedexLocationAreaEP
	}

	res, err := http.Get(areaName)

	if err != nil {
		return Page{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return Page{}, err
	}

	cache.Add(areaName, data)

	err = json.Unmarshal(data, &page)

	if err != nil {
		return Page{}, err
	}

	return page, nil
}

func (c *Client) GetLocationArea(areaName string, cache *pokecahe.Cache) (area LocationArea, err error) {

	if len(areaName) == 0 {
		return LocationArea{}, errors.New("poke_api: areaName is empty")
	}

	areaURL := pokedexLocationAreaEP + areaName
	if entry, is := cache.Get(areaURL); is {
		err = json.Unmarshal(entry, &area)

		if err != nil {
			return LocationArea{}, err
		}

		return area, nil
	}

	req, err := http.NewRequest("GET", areaURL, nil)

	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return LocationArea{}, fmt.Errorf("poke_api: status %d: %s", resp.StatusCode, body)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationArea{}, err
	}

	err = json.Unmarshal(data, &area)

	if err != nil {
		return LocationArea{}, err
	}

	cache.Add(areaURL, data)

	return area, nil
}

func (client *Client) GetPokemon(pokemonName string, cache *pokecahe.Cache) (pokemon Pokemon, err error) {
	if len(pokemonName) == 0 {
		return Pokemon{}, errors.New("poke_api:GetPokemon() pokemonName is empty\n")
	}

	pokemonURl := pokedexPokemonEP + pokemonName
	if entry, is := cache.Get(pokemonURl); is {
		err = json.Unmarshal(entry, &pokemon)

		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", pokemonURl, nil)

	if err != nil {
		return Pokemon{}, err
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		if resp.StatusCode == http.StatusNotFound {
			return Pokemon{}, fmt.Errorf("No pokemon named %s found, try again\n", pokemonName)
		} else {
			return Pokemon{}, fmt.Errorf("poke_api:GetPokemon() status %d: %s\n", resp.StatusCode, body)
		}
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	cache.Add(pokemonURl, data)

	return pokemon, nil
}
