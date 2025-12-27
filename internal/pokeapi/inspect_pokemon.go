package pokeapi


import (
    "fmt"
)


func (c *Client) InspectPokemon(name string) error {
    pokemon, err := c.pokedex.Get(name)

    if err != nil { 
        return err
    }

    fmt.Printf("Name: %s", pokemon.Name)

    return nil
}
