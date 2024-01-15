package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)

func (c *Client) ListPokemonEncounters(locationAreaName string) (LocationEncounterResponse, error) {
	var dex LocationEncounterResponse

	fullURL := baseURL + "/location-area/" + locationAreaName

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationEncounterResponse{}, err
	}

	// Check the cache
	dat,ok := c.cache.Get(fullURL)
	if ok {
		err = json.Unmarshal(dat, &dex)
		if err != nil {
			return LocationEncounterResponse{}, err
		}
		return  dex, nil
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationEncounterResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationEncounterResponse{}, fmt.Errorf("Error: %d", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationEncounterResponse{}, err
	}

	err = json.Unmarshal(dat, &dex)
	if err != nil {
		return LocationEncounterResponse{}, err
	}

	c.cache.Add(fullURL, dat)
	return  dex, nil
}