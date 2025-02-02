package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonData(pokemon *string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + *pokemon

	if val, ok := c.cache.Get(url); ok {
		pokemonData := Pokemon{}
		err := json.Unmarshal(val, &pokemonData)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonData, nil
	}

	//Create new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	//Do request to get response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	//Read response to get data stream
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	//Unmarshal the data to go struct
	pokemonData := Pokemon{}
	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonData, nil
}
