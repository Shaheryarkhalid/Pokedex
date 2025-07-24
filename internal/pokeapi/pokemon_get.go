package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetPokemon(pokemonName string) (pokemon Pokemon, err error) {
	pokemonName = strings.Trim(pokemonName, " ")
	if pokemonName == "" {
		return pokemon, fmt.Errorf("Pokemon Name cannot be empty.")
	}
	apiUrl := baseURL + "/pokemon/" + pokemonName
	data, found := c.cache.Get(apiUrl)
	if !found {
		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			return pokemon, fmt.Errorf("Error creating the Api request: %v", err)
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return pokemon, fmt.Errorf("Error making the Api request: %v", err)
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return pokemon, fmt.Errorf("Error reading the Api request: %v", err)
		}
		c.cache.Add(apiUrl, data)
	}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		fmt.Printf("Error:%v %v", pokemonName, string(data))
		fmt.Println()
		return pokemon, nil
	}
	return pokemon, err
}
