package pokeapi


import( 
    "fmt"
)


type Pokedex map[string]Pokemon

func NewPokedex() *Pokedex {
    pokedex := make(Pokedex)
    return &pokedex
}

func (p Pokedex) Add(pokemon Pokemon) {
    p[pokemon.Name] = pokemon
}

func (p Pokedex) Get(name string) (Pokemon, error) { 
    pokemon, exists := p[name]
    if !exists {
        return Pokemon{}, fmt.Errorf("you have not caught that pokemon")
    }

    return pokemon, nil
}

func (p Pokedex) List() error {
    if len(p) == 0 { 
        return fmt.Errorf("You haven't captured any Pokemon!")
    }

    fmt.Println("Your Pokemon:")
    for _, pokemon := range p { 
        fmt.Printf(" - %s\n", pokemon.Name)
    }

    return nil
}
