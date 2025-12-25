package pokedex


import( 
    "fmt"
)


type Pokemon struct {
    name string
}

type Pokedex map[string]Pokemon

func NewPokemon(name string) Pokemon { 
    return Pokemon {
        name: name,
    }
}

func NewPokedex() *Pokedex {
    pokedex := make(Pokedex)
    return &pokedex
}

func (p Pokedex) Add(name string) {
    p[name] = NewPokemon(name)
    p.List()
}

func (p Pokedex) List() {
    if len(p) == 0 { 
        fmt.Println("You haven't captured any Pokemon!")
    }

    fmt.Println("You've captured the following Pokemon:")
    for _, pokemon := range p { 
        fmt.Printf(" - %s\n", pokemon.name)
    }
}
