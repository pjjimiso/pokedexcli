package main


import (
    "fmt"
    "errors"
)


func commandMapb(cfg *config, args ...string) error {
    if cfg.prevLocationsURL == nil || *cfg.prevLocationsURL == "" {
        return errors.New("you're on the first page")
    }

    laMap, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationsURL)
    if err != nil { 
        return err
    }

    cfg.nextLocationsURL = laMap.Next
    cfg.prevLocationsURL = laMap.Previous

    for _, location := range laMap.Results { 
        fmt.Println(location.Name)
    }

    return nil
}

