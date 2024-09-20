package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\n\nUsage:\n\n")

	availableCommands := getCommands()

	for _, command := range availableCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Print("\n\n")
	return nil
}
