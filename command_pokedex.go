package main


func commandPokedex(cfg *config, args ...string) error {
    err := cfg.pokeapiClient.ListPokemon()
    if err != nil { 
        return err
    }

    return nil
}
