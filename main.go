package main

import (
	"github.com/aecef/pokedexcli/internal/pokeapi"
	"github.com/aecef/pokedexcli/internal/pokebag"
	"time"
)

type config struct {
	pokeapiClient	pokeapi.Client
	pokebag			pokebag.Pokebag
	nextLocationURL *string
	prevLocationURL *string
}


func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokebag: pokebag.NewPokebag(),
	}

	startRepl(&cfg)
}
