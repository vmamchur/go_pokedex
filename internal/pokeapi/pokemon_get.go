package pokeapi

import (
    "encoding/json"
    "io"
    "net/http"
)

func (c *Client) GetPokemon(name string) (RespShallowPokemon, error) {
	url := baseURL + "/pokemon/" + name

    if cachedData, ok := c.cache.Get(url); ok {
        pokemonResp := RespShallowPokemon{}
        err := json.Unmarshal(cachedData, &pokemonResp)
        if err != nil {
            return RespShallowPokemon{}, err
        }
        return pokemonResp, nil
    }
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return RespShallowPokemon{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return RespShallowPokemon{}, err
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return RespShallowPokemon{}, err
    }

    pokemonResp := RespShallowPokemon{}
    err = json.Unmarshal(data, &pokemonResp)
    if err != nil {
        return RespShallowPokemon{}, err
    }

    c.cache.Add(url, data)

    return pokemonResp, nil
}
