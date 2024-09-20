package main

import (
	"time"

	"github.com/mirandajoy/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(time.Hour, 5*time.Second),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
