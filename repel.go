package main

import (
	"Pokedex/internal/pokeapi"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	pokeapiClient       pokeapi.Client
	nextLocationUrl     *string
	previousLocationUrl *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func startRepel(config *Config) {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !s.Scan() {
			break
		}
		userInput := cleanInput(s.Text())
		if len(userInput) == 0 {
			continue
		}
		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}
		c, exists := getCommands()[userInput[0]]
		if !exists {
			fmt.Println("Unknown Command.")
			continue
		}
		err := c.callback(args, config)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

type cliCommands struct {
	name        string
	description string
	callback    func([]string, *Config) error
}

func cleanInput(input string) (cleanedInput []string) {
	cleanedInput = strings.Split(strings.ToLower(strings.Trim(input, " ")), " ")
	for idx, word := range cleanedInput {
		cleanedInput[idx] = strings.Trim(word, " ")
	}
	return cleanedInput
}
func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"map": {
			name:        "map",
			description: "Gets list of Location Areas from Pokemon API.\n",
			callback:    cmdMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets list of Previously viewed Location Areas from Pokemon API.\n",
			callback:    cmdMapb,
		},
		"explore": {
			name:        "explore",
			description: "Gets the list of pokemons in a given area. \n \t - Takes location area name as an argument.\n",
			callback:    cmdExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the given pokemon.\n \t - Takes Pokemon name as an argument.",
			callback:    cmdCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays details of the given caught pokemon.",
			callback:    cmdInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "Displays all of the caught Pokemon.",
			callback:    cmdPokedex,
		},

		"help": {
			name:        "help",
			description: "Displays all of the available commands.\n",
			callback:    cmdHelp,
		},
		"exit": {
			name:        "exit",
			description: "Closes the program.\n",
			callback:    cmdExit,
		},
		"clear": {
			name:        "clear",
			description: "Clears the terminal screen.\n",
			callback:    cmdClear,
		},
	}

}
