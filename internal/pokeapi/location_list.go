package pokeapi


import (
    "encoding/json"
    "io"
    "net/http"
)


func (c *Client) ListLocationAreas(pageURL *string) (locationAreaMap, error) {
    url := baseURL + "/location-area"

    if pageURL != nil { 
        url = *pageURL
    }

    var err error

    jsonData, exists := c.cache.Get(url)
    if !exists {
        jsonData, err = c.getLocationsFromHTTP(url)
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

func (c *Client) getLocationsFromHTTP(url string) ([]byte, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil { 
        return []byte{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil { 
        return []byte{}, err
    }
    defer res.Body.Close()
    
    jsonData, err := io.ReadAll(res.Body)
    if err != nil { 
        return []byte{}, err
    }

    return jsonData, nil
}
