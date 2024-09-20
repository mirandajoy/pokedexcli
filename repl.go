package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("==POKEDEX==\n\n")

	for {
		fmt.Print("Pokedex >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		if len(text) == 0 {
			continue
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.callback(cfg, args...)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Provide 20 new locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Provide previous 20 locations",
			callback:    commandBMap,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "See pokemon in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon}",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View my pokemon",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
