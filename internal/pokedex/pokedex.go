package pokedex

type Pokemon struct {
    name string
}

type Pokedex map[string]Pokemon

func NewPokedex() *Pokedex {
    pokedex := make(Pokedex)
    return &pokedex
}

func (p Pokedex) Add(pokemon Pokemon) {
    p[pokemon.name] = pokemon
}
