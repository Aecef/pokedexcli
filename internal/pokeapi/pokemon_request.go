package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)

func (c *Client) GetPokemon(name string) (PokemonResponse, error) {
	var pokemon PokemonResponse

	fullURL := baseURL + "/pokemon/" + name

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResponse{}, err
	}
	
	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit!")
		err = json.Unmarshal(dat, &pokemon)
		if err != nil {
			return PokemonResponse{}, err
		}
		return pokemon, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonResponse{}, fmt.Errorf("Error: %d", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(fullURL, dat)
	return pokemon, nil
}