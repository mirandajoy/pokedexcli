package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	fmt.Println("Exploring...")
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationAreas(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:", locationAreaName)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
