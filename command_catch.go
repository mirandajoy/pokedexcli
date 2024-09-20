package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	fmt.Println("Catching...")
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)

	if randNum > threshold {
		fmt.Printf("You failed to catch %s", pokemon.Name)
		return nil
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("You caught a %s\n\n", pokemon.Name)

	return nil
}
