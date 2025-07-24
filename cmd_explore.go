package main

import "fmt"

func cmdExplore(args []string, config *Config) error {
	if len(args) < 1 || args[0] == "" {
		fmt.Println("Location Area name must be provided as an argument.")
		return nil
	}
	locationArea, err := config.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}
	for _, pokemonEncounters := range locationArea.PokemonEncounters {
		fmt.Println(pokemonEncounters.Pokemon.Name)
	}
	return nil
}
