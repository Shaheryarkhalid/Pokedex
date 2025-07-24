package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocationAreas(url *string) (locationAreas ShallowLocation, err error) {
	apiUrl := baseURL + "/location-area"
	if url != nil {
		apiUrl = *url
	}

	data, found := c.cache.Get(apiUrl)
	if !found {
		httpClient := c.httpClient
		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			return locationAreas, fmt.Errorf("Error creating request for Location Area: %v", err)
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			return locationAreas, fmt.Errorf("Error making request for Location Area: %v", err)
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return locationAreas, fmt.Errorf("Error Reading the response body: %v", err)
		}
		c.cache.Add(apiUrl, data)
	}

	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		fmt.Printf("Error: %v", string(data))
		fmt.Println()
		return locationAreas, nil
	}
	return locationAreas, nil
}
