package main

import (
	"os"
	"fmt"
	"bufio"
	"errors"
	"strings"
)

func startRepl(cfg *config) {
	
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		
		cleaned := cleanInput(text)
		commandName := cleaned[0]

		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			errors.New("Unknown command:" + commandName)
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print("pokedex > ")

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

}


func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words
}