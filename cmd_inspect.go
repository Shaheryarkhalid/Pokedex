package main

import (
	"fmt"
	"strings"
)

func cmdInspect(args []string, config *Config) error {
	if len(args) == 0 || strings.Trim(args[0], " ") == "" {
		fmt.Println("Pokemon name must be provided.")
		return nil
	}
	pokemon, found := config.caughtPokemon[args[0]]
	if !found {
		fmt.Printf("%v has not be caught yet.\n", args[0])
		return nil
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
