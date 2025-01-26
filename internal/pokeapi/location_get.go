package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (RespShallowLocation, error) {
	url := baseURL + "/location-area/" + name

    if cachedData, ok := c.cache.Get(url); ok {
        locationResp := RespShallowLocation{}
        err := json.Unmarshal(cachedData, &locationResp)
        if err != nil {
            return RespShallowLocation{}, err
        }
        return locationResp, nil
    }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}

	locationResp := RespShallowLocation{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocation{}, err
	}

    c.cache.Add(url, data)

	return locationResp, nil
}
