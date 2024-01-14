package main

import (
	"os"
	"fmt"
	"log"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
	return commands
}

func commandHelp(cfg *config) error {
	fmt.Println("Available commands:")
	fmt.Println("help - Shows this help message")
	fmt.Println("exit - Exits the pokedex")
	return nil
}
func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		log.Fatal(err)
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

func commandMapb(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationURL)
	if err != nil {
		log.Fatal(err)
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
