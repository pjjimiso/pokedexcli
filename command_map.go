package main


import (
    "fmt"
    "errors"
)


func commandMap(cfg *config) error { 
    laMap, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationsURL)
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

func commandMapb(cfg *config) error {
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
