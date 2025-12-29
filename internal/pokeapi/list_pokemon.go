package pokeapi


func (c *Client) ListPokemon() error {
    err := c.pokedex.List()

    if err != nil { 
        return err
    }

    return nil
}
