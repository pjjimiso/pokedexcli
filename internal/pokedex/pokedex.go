package pokedex

type Pokemon struct {
    name string
}

type Pokedex struct { 
    map[string]Pokemon
}

func NewPokedex() Pokedex {
    return Pokedex {
        make(map[string]Pokemon)
    }
}

func (p Pokedex) Add(pokemon Pokemon) {
    p[pokemon.Name] = pokemon
}
