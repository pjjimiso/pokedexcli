package main


import(
    "fmt"
)


func commandCatch(cfg *config, args ...string) error {
    if len(args) == 0 {
        return fmt.Errorf("specify a pokemon to catch...")
    }
    pokemon := args[0]

    err := cfg.pokeapiClient.CatchPokemon(pokemon)
    if err != nil { 
        return err
    }

    return nil
}
