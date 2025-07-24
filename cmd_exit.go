package main

import (
	"fmt"
	"os"
)

func cmdExit(args []string, config *Config) error {
	fmt.Println("Closing the Pokedex. Bye...")
	os.Exit(0)
	return nil
}
