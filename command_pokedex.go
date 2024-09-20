package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokemonList := cfg.caughtPokemon
	if len(pokemonList) == 0 {
		return fmt.Errorf("you haven't caught any pokeman")
	}

	fmt.Printf("Your Pokemon:\n")
	for _, pokeman := range pokemonList {
		fmt.Printf("- %s\n", pokeman.Name)
	}

	return nil
}
