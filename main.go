package main


import (
    "time"

    "github.com/pjjimiso/pokedexcli/internal/pokeapi"
    "github.com/pjjimiso/pokedexcli/internal/pokecache"
)


func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second) 
    pokeCache := pokecache.NewCache(5 * time.Second)
    cfg := &config{
        pokeapiClient: pokeClient,
        pokeapiCache: pokeCache,
    }

    startRepl(cfg)
}
