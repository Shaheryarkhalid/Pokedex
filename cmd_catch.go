package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func cmdCatch(args []string, config *Config) error {
	if len(args) < 1 || strings.Trim(args[0], " ") == "" {
		fmt.Println("Pokemon name must be provided.")
		return nil
	}
	if _, ok := config.caughtPokemon[args[0]]; ok {
		fmt.Printf("%v has already been caught.\n", args[0])
		return nil
	}
	pokemon, err := config.pokeapiClient.GetPokemon(args[0])
	if err == nil && pokemon.Name == "" {
		return nil
	}
	if err != nil {
		return err
	}
	chance := 1.0 - float64(pokemon.BaseExperience)/366.0
	myChance := rand.Float64()
	if myChance > chance {
		config.caughtPokemon[pokemon.Name] = pokemon
		fmt.Printf("%v was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
	return nil
}
