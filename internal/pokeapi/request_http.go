package pokeapi


import (
    "io"
    "net/http"
    "fmt"
)


func (c *Client) requestHTTP(url string) ([]byte, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil { 
        return []byte{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil { 
        return []byte{}, err
    }
    defer res.Body.Close()

    if res.StatusCode > 299 {
        return []byte{}, fmt.Errorf("http request returned status code %d", res.StatusCode)
    }

    jsonData, err := io.ReadAll(res.Body)
    if err != nil { 
        return []byte{}, err
    }

    return jsonData, nil
}
