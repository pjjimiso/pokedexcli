package pokeapi


import (
    "fmt"
)


func (c *Client) InspectPokemon(name string) error {
    pokemon, err := c.pokedex.Get(name)

    if err != nil { 
        return err
    }

    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)

    fmt.Printf("Stats:\n")
    for _, stat := range pokemon.Stats {
        fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Printf("Types:\n")
    for _, t := range pokemon.Types {
        fmt.Printf("  - %s\n", t.Type.Name)
    }

    return nil
}
