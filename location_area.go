package main


type locationAreaMap struct { 
    Count       int             `json:"count"`
    Next        string          `json:"next"`
    Previous    string          `json:"previous"`
    Results     []struct {
        Name    string  `json:"name"`
        Url     string  `json:"url"`
    }  `json:"results"`
}

