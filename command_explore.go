package main


import (
    "fmt" 
    "errors"
)


func commandExplore(cfg *config, location string) error { 
    pokemon, err := pokeapiClient.ListPokemon(location)
    if err != nil { 
        return err
    }

    return nil
}
