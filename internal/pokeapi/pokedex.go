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
    p.List()
}

func (p Pokedex) Get(name string) (Pokemon, error) { 
    pokemon, exists := p[name]
    if !exists {
        return Pokemon{}, fmt.Errorf("you have not caught that Pokemon")
    }

    return pokemon, nil
}

func (p Pokedex) List() {
    if len(p) == 0 { 
        fmt.Println("You haven't captured any Pokemon!")
    }

    fmt.Println("You've captured the following Pokemon:")
    for _, pokemon := range p { 
        fmt.Printf(" - %s\n", pokemon.Name)
    }
}
