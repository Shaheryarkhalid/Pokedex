package main

import "fmt"

func cmdMapf(args []string, config *Config) error {
	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.nextLocationUrl)
	if err != nil {
		return err
	}
	config.previousLocationUrl = config.nextLocationUrl
	config.nextLocationUrl = locationAreas.Next
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func cmdMapb(args []string, config *Config) error {
	if config.previousLocationUrl == nil {
		return fmt.Errorf("You are on first page.")
	}
	locationAreas, err := config.pokeapiClient.GetLocationAreas(config.previousLocationUrl)
	if err != nil {
		return err
	}
	config.nextLocationUrl = locationAreas.Next
	config.previousLocationUrl = locationAreas.Previous
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}
