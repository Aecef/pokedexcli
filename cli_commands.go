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
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Displays information about the selected pokemon",
			callback:    commandInspect,
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
	catchBase:= rand.Intn(baseXP)
	fBaseXP := float64(baseXP)
	catchdDifficulty := 0.5 + (0.15 * (fBaseXP/100))
	catchChance := float64(catchBase) / fBaseXP

	fmt.Printf("Throwing a pokeball at %s...\n", pokemon.Name)
	fmt.Printf("%.2f and you needed at least %.2f\n", catchChance, catchdDifficulty)
	if catchChance >= catchdDifficulty ||  catchChance >= 0.99   {
		fmt.Printf("You caught %s!\n", pokemon.Name)
		cfg.pokebag.Add(pokemon)
	} else {
		fmt.Printf("%s broke free!\n", pokemon.Name)
	}
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	targetPokemon := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(targetPokemon)
	if err != nil {
		return err
	}

	_, ok := cfg.pokebag.Get(pokemon.Name)
	if !ok {
		fmt.Printf("You have not caught %s!\n", pokemon.Name)
		return nil
	}


	fmt.Println("=======================================")
	fmt.Printf("Pokemon: %s\n", pokemon.Name)
	fmt.Println("=======================================")
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	fmt.Println("=======================================")

	return nil
}
