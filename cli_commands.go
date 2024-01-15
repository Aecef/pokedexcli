package main

import (
	"os"
	"fmt"
	"errors"
	"math/rand"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Shows a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas from the pokeapi",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas from the pokeapi",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_area>",
			description: "Displays possible pokemon to catch in the selected location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempts to catch the selected pokemon",
			callback:    commandCatch,
		},
	}
	return commands
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Available commands:")

	for _, command := range getCommands() {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}

	return nil
}
func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	fmt.Println("=======================================")
	fmt.Println("LocationAreas: ")
	fmt.Println("=======================================")
	for _, locationArea := range resp.Results {
		fmt.Println(" - " + locationArea.Name)
	}
	fmt.Println("=======================================")
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("Youre on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationURL)
	if err != nil {
		return err
	}
	fmt.Println("=======================================")
	fmt.Println("LocationAreas: ")
	fmt.Println("=======================================")
	for _, locationArea := range resp.Results {
		fmt.Println(" - " + locationArea.Name)
	}
	fmt.Println("=======================================")
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("You must provide a location area")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.ListPokemonEncounters(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Println("=======================================")
	fmt.Printf("Pokemon Encountered in %s: ", locationArea.Name)
	fmt.Println("\n=======================================")
	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		fmt.Println(" - " + pokemonEncounter.Pokemon.Name)
	}
	fmt.Println("=======================================")

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	targetPokemon := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(targetPokemon)
	if err != nil {
		return err
	}

	_, ok := cfg.pokebag.Get(pokemon.Name)
	if ok {
		fmt.Printf("%s is already in your pokebag!\n", pokemon.Name)
		return nil
	}

	baseXP := pokemon.BaseExperience
	catchChance := rand.Intn(baseXP)
	fmt.Printf("Throwing a pokeball at %s...\n", pokemon.Name)
	if float64(catchChance) > float64(baseXP) / 2  {
		fmt.Printf("You caught %s!\n", pokemon.Name)
		cfg.pokebag.Add(pokemon)
	} else {
		fmt.Printf("%s broke free!\n", pokemon.Name)
	}
	return nil

}
