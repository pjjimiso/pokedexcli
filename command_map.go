package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
)


func getMap(url string) (locationAreaMap, error) {
    res, err := http.Get(url)
    if err != nil { 
        return locationAreaMap{}, err
    }
    defer res.Body.Close()

    jsonData, err := io.ReadAll(res.Body)
    if err != nil { 
        return locationAreaMap{}, err
    }

    locationMap := locationAreaMap{}
    err = json.Unmarshal(jsonData, &locationMap)
    if err != nil { 
        return locationAreaMap{}, err
    }

    config.Next = locationMap.Next
    config.Previous = locationMap.Previous

    return locationMap, nil
}

func printMap(laMap locationAreaMap) {
    fmt.Println("Next page:", laMap.Next) 
    fmt.Println("Previous page:", laMap.Previous) 

    for _, location := range laMap.Results { 
        fmt.Println(location.Name)
    }
}

func commandMap() error { 
    laMap, err := getMap(config.Next) 
    if err != nil { 
        return err
    }

    printMap(laMap) 

    return nil
}

func commandMapb() error {
    if config.Previous == "" {
        fmt.Println("you're on the first page")
        return nil
    }

    laMap, err := getMap(config.Previous)
    if err != nil { 
        return err
    }

    printMap(laMap)

    return nil
}
