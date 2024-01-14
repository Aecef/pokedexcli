package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

func (c *Client) ListLocationAreas() (LocationAreasResponse, error) {
	var dex LocationAreasResponse
	fullURL := baseURL + "/location-area"
	
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("Error: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	err = json.Unmarshal(data, &dex)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return  dex, nil
}