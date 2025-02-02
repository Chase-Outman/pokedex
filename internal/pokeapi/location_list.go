package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) DeepListLocation(area *string) (RespDeepLocation, error) {
	url := baseURL + *area

	if val, ok := c.cache.Get(url); ok {
		deepLocatoinResp := RespDeepLocation{}
		err := json.Unmarshal(val, &deepLocatoinResp)
		if err != nil {
			return RespDeepLocation{}, err
		}
		return deepLocatoinResp, nil
	}

	//Create new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocation{}, err
	}
	//Do request to get response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocation{}, err
	}
	defer resp.Body.Close()
	//Read response to get data stream
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocation{}, err
	}
	//Unmarshal the data to go struct
	deepLocationResp := RespDeepLocation{}
	err = json.Unmarshal(data, &deepLocationResp)
	if err != nil {
		return RespDeepLocation{}, err
	}
	c.cache.Add(url, data)
	return deepLocationResp, nil

}

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
