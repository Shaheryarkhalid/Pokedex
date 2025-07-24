package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetLocation(areaName string) (locationArea LocationArea, err error) {
	areaName = strings.Trim(areaName, " ")
	if areaName == "" {
		return locationArea, fmt.Errorf("Area name must be provided.")
	}
	apiUrl := baseURL + "/location-area/" + areaName
	data, found := c.cache.Get(apiUrl)
	if !found {
		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			return locationArea, fmt.Errorf("Error creating Api request: %v", err)
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return locationArea, fmt.Errorf("Error making Api request: %v", err)
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return locationArea, fmt.Errorf("Error Api response body: %v", err)
		}
		c.cache.Add(apiUrl, data)
	}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		fmt.Printf("Error: %v", string(data))
		fmt.Println()
		return locationArea, nil
	}
	return locationArea, err
}
