package pokebag

import (
	"github.com/aecef/pokedexcli/internal/pokeapi"
)

type Pokebag struct {
	pokemon map[string]pokeapi.PokemonResponse
}

func NewPokebag() Pokebag {
	bag := Pokebag{
		pokemon: make(map[string]pokeapi.PokemonResponse),
	}
	return bag 
}

func (b *Pokebag) Add(pokemon pokeapi.PokemonResponse) {
	b.pokemon[pokemon.Name] = pokemon
}

func (b *Pokebag) Get(name string) (pokeapi.PokemonResponse, bool) {
	pokemon, ok := b.pokemon[name]
	if !ok {
		return pokeapi.PokemonResponse{}, false
	}
	return pokemon, ok
}

func (b *Pokebag) AllCaughtPokemon() []string {
	caught := []string{}
	for _, pokemon := range b.pokemon {
		caught = append(caught, pokemon.Name)
	}
	return caught
}