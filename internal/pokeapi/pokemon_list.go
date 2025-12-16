package pokeapi


import (
    "encoding/json"
    "fmt"
)


func (c *Client) ListPokemon(location string) (error) {
    url := baseURL + "/location-area/" + location

    var err error

    jsonData, exists := c.cache.Get(url)
    if !exists {
        jsonData, err = c.requestHTTP(url)
        if err != nil {
            return err
        }
        c.cache.Add(url, jsonData)
    }

    locationData := locationData{}
    err = json.Unmarshal(jsonData, &locationData)
    if err != nil { 
        return err
    }

    fmt.Println("locationData: ", locationData)

    return nil
}

