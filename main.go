package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

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

func main() {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

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
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		command, ok := commands[text]
		if !ok {
			errors.New("Unknown command:" + text)
			continue
		}
		command.callback()
		fmt.Print("pokedex > ")

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

}
