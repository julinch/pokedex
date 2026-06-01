package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	pokecahe "github.com/julinch/pokedex/internal/pokecache"
)

const pokedexEP = "https://pokeapi.co/api/v2/location-area/"

func GetPage(areaName string, cache *pokecahe.Cache) (page Page, err error) {

	if entry, is := cache.Get(areaName); is {
		err = json.Unmarshal(entry, &page)

		if err != nil {
			return Page{}, err
		}

		return page, nil
	}

	if len(areaName) == 0 {
		areaName = pokedexEP
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

	areaURL := pokedexEP + areaName
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
