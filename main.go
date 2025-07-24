package main

import (
	"Pokedex/internal/pokeapi"
	"time"
)

func main() {
	con := Config{
		pokeapiClient: pokeapi.NewClient(10*time.Second, 5*time.Minute),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepel(&con)

}
