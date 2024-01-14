package pokeapi

import (
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client {
			Timeout: time.Minute * 10,
		},
	}
}

var dex LocationAreasResponse

func CommandMap() error {
	fmt.Println("Map:")
	fmt.Println("Pallet Town")
	fmt.Println("Route 1")
	return nil
}

func CommandMapb() error {
	fmt.Println("Map:")
	fmt.Println("Pallet Town")
	fmt.Println("Route 1")
	return nil
}