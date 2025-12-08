package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
)


func getMap() (locationAreaMap, error) {

    res, err := http.Get(config.Next)
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

    laMap, err := getMap() 
    if err != nil { 
        return err
    }
    
    config.Next = laMap.Next
    config.Previous = laMap.Previous

    printMap(laMap) 

    return nil
}
