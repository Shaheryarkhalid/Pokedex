package main

import "fmt"

func cmdPokedex(args []string, config *Config) error {
	if len(config.caughtPokemon) == 0 {
		fmt.Println("No Pokemon has been caught yet.")
		return nil
	}
	for _, pokemon := range config.caughtPokemon {
		fmt.Printf("\t-%v\n", pokemon.Name)
	}
	return nil
}
