package pokeapi


import (
    "encoding/json"
)


func (c *Client) ListLocationAreas(pageURL *string) (locationAreaMap, error) {
    url := baseURL + "/location-area"

    if pageURL != nil { 
        url = *pageURL
    }

    var err error

    jsonData, exists := c.cache.Get(url)
    if !exists {
        jsonData, err = c.requestHTTP(url)
        if err != nil { 
            return locationAreaMap{}, err
        }
        c.cache.Add(url, jsonData)
    }

    locationMap := locationAreaMap{}
    err = json.Unmarshal(jsonData, &locationMap)
    if err != nil { 
        return locationAreaMap{}, err
    }

    return locationMap, nil
}

