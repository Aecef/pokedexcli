package main

import (
	"github.com/aecef/pokedexcli/internal/pokeapi"
	"os"
	"fmt"
	"log"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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

func commandHelp() error {
	fmt.Println("Available commands:")
	fmt.Println("help - Shows this help message")
	fmt.Println("exit - Exits the pokedex")
	return nil
}
func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
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

	return nil
}



func commandMapb() error {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=======================================")
	fmt.Println("LocationAreas: ")
	fmt.Println("=======================================")
	for _, locationArea := range resp.Results {
		fmt.Println(locationArea.Name)
	}
	fmt.Println("=======================================")

	return nil
}