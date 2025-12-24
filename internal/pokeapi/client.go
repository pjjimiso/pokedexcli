package pokeapi

import ( 
    "net/http"
    "time"

    "github.com/pjjimiso/pokedexcli/internal/pokecache"
    "github.com/pjjimiso/pokedexcli/internal/pokedex"
)


type Client struct { 
    httpClient  http.Client
    cache       *pokecache.Cache
    pokedex     *pokedex.Pokedex
}

func NewClient(timeout, cacheInterval time.Duration) Client { 
    return Client { 
        httpClient: http.Client{
            Timeout: timeout,
        },
        cache: pokecache.NewCache(cacheInterval),
        pokedex: pokedex.NewPokedex(),
    }
}
