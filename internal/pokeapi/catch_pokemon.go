package pokeapi


import (
    "encoding/json"
    "fmt"
    "math/rand"

    "github.com/pjjimiso/pokedexcli/internal/pokedex"
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

    baseExp := pokemonStats.BaseExperience
    scalingFactor := 200.0
    catchChance := 1.0 / (1.0 + float64(baseExp) / scalingFactor)

    fmt.Printf("%s base_experience is %d, your chance to catch is %.2f\n", pokemon, baseExp, catchChance)

    if rand.Float64() < catchChance {
        fmt.Printf("%s was caught!\n", pokemon)
        // TODO - create Pokemon to pass in
        c.pokedex.Add(pokedex.Pokemon{})
    } else {
        fmt.Printf("%s escaped!\n", pokemon)
    }

    return nil
}

