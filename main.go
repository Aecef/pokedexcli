package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("pokedex > ")

	// Read input line by line
	for scanner.Scan() {
		text := scanner.Text() // Get the current line of text
		if text == "" {
			break // Exit loop if an empty line is entered
		}
		fmt.Println("You entered:", text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

}
