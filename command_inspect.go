package main


import(
    "fmt"
)


func commandInspect(cfg *config, args ...string) error {
    if len(args) == 0 {
        return fmt.Errorf("specify a pokemon to inspect...")
    }
    pokemon := args[0]

    err := cfg.pokeapiClient.InspectPokemon(pokemon)
    if err != nil { 
        return err
    }

    return nil
}
