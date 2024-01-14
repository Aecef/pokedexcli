package main

import (
	"github.com/aecef/pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient	pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&cfg)
}
