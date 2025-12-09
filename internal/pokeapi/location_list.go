package pokeapi


import (
    "encoding/json"
    "io"
    "net/http"
)


func (c *Client) GetLocationAreas(pageURL *string) (locationAreaMap, error) {
    url := baseURL + "/location-area"

    if pageURL != nil { 
        url = *pageURL
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil { 
        return locationAreaMap{}, err
    }

    res, err := c.httpClient.Do(req)
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
