package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	
	var dex LocationAreasResponse
	fullURL := baseURL + "/location-area"
	
	if pageURL != nil {
		fullURL = *pageURL
	}
	
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	// Check the cache
	dat,ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit")
		err = json.Unmarshal(dat, &dex)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return  dex, nil
	}
	fmt.Println("Cache miss")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("Error: %d", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	err = json.Unmarshal(dat, &dex)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, dat)
	return  dex, nil
}