package pokeapi


import (
    "encoding/json"
    "fmt"
)


func (c *Client) ListPokemon(location string) ([]string, error) {
    url := baseURL + "/location-area/" + location

    var err error

    jsonData, exists := c.cache.Get(url)
    if !exists {
        jsonData, err = c.requestHTTP(url)
        if err != nil {
            return nil, err
        }
        c.cache.Add(url, jsonData)
    }

    fmt.Printf("Exploring %s...\n", location)

    locationData := locationData{}
    err = json.Unmarshal(jsonData, &locationData)
    if err != nil { 
        return nil, err
    }

    fmt.Println("Found Pokemon:")

    for _, encounter := range locationData.PokemonEncounters {
        fmt.Println(" -", encounter.Pokemon.Name)
    }

    //pokemonFound := []string{}

    return []string{}, nil
}

