package pokeapi


import (
    "encoding/json"
    "fmt"
    "math/rand"
)


func (c *Client) CatchPokemon(name string) error { 
    url := baseURL + "/pokemon/" + name

    var err error

    jsonData, exists := c.cache.Get(url)
    if ! exists {
        jsonData, err = c.requestHTTP(url)
        if err != nil {
            return err
        }
        c.cache.Add(url, jsonData)
    }

    pokemon := Pokemon{} 
    err = json.Unmarshal(jsonData, &pokemon)
    if err != nil { 
        return err
    }

    fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

    baseExp := pokemon.BaseExperience
    // TODO - make this a global config
    scalingFactor := 200.0
    catchChance := 1.0 / (1.0 + float64(baseExp) / scalingFactor)

    fmt.Printf("%s base_experience is %d, your chance to catch is %.2f\n", pokemon.Name, baseExp, catchChance)

    if rand.Float64() < catchChance {
        fmt.Printf("%s was caught!\n", pokemon.Name)
        c.pokedex.Add(pokemon)
    } else {
        fmt.Printf("%s escaped!\n", pokemon.Name)
    }

    return nil
}

