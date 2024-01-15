package main

import (
	"os"
	"fmt"
	"bufio"
	"errors"
)


func startRepl(cfg *config) {
	
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		command, ok := getCommands()[text]
		if !ok {
			errors.New("Unknown command:" + text)
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print("pokedex > ")

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

}