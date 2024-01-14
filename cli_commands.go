package main

import (
	"github.com/aecef/pokedexcli/internal/pokeapi"
	"os"
	"fmt"
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
			description: "Shows the map",
			callback:    pokeapi.CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the map",
			callback:    pokeapi.CommandMapb,
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