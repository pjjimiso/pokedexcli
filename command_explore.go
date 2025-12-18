package main


import (
    "fmt"
)


func commandExplore(cfg *config, args ...string) error { 
    if len(args) == 0 {
        return fmt.Errorf("explore requires a location argument")
    }
    location := args[0]

    err := cfg.pokeapiClient.ListPokemon(location)
    if err != nil { 
        return err
    }

    return nil
}
