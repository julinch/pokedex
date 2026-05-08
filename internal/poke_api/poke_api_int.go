package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	pokecahe "pokedex/internal/pokecache"
)

const pokedexEP = "https://pokeapi.co/api/v2/location-area/"

func GetPage(url string, cache *pokecahe.Cache) (page Page, err error) {

	if entry, is := cache.Get(url); is {
		err = json.Unmarshal(entry, &page)

		if err != nil {
			return Page{}, err
		}

		return page, nil
	}

	if len(url) == 0 {
		url = pokedexEP
	}

	res, err := http.Get(url)

	if err != nil {
		return Page{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return Page{}, err
	}

	cache.Add(url, data)

	err = json.Unmarshal(data, &page)

	if err != nil {
		return Page{}, err
	}

	return page, nil
}
