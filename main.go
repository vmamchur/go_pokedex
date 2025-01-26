package main

import (
	"time"

	"github.com/vmamchur/go_pokedex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeapiClient,
		caughtPokemon: map[string]pokeapi.RespShallowPokemon{},
	}

	startRepl(cfg)
}
