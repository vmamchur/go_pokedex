package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemonResp.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonResp.Name)

    cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
	return nil
}
