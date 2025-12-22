package pokeapi


import (
    "encoding/json"
    "fmt"
)


func (c *Client) CatchPokemon(pokemon string) error { 
    url := baseURL + "/pokemon/" + pokemon

    var err error

    jsonData, exists := c.cache.Get(url)
    if ! exists {
        jsonData, err = c.requestHTTP(url)
        if err != nil {
            return err
        }
        c.cache.Add(url, jsonData)
    }

    pokemonStats := pokemonStats{} 
    err = json.Unmarshal(jsonData, &pokemonStats)
    if err != nil { 
        return err
    }

    fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
    
    // attempt to catch pokemon

    return nil
}

