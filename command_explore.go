package main

import (
	"fmt"
    "errors"
)

func commandExplore(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("you must provide a location name")
    }

    locationName := args[0]
	locationResp, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

    fmt.Printf("Exploring %s...\n", locationName)
    fmt.Println("Found Pokemon:")
	for _, enc := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
