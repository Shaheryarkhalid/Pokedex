package main

import "fmt"

func cmdHelp(args []string, config *Config) error {
	fmt.Println("\nWelcome to Pokedex.")
	fmt.Println("Here are some commands to get started.")
	fmt.Println()
	for _, val := range getCommands() {
		fmt.Printf("%v : %v", val.name, val.description)
		fmt.Println()
	}
	return nil
}
