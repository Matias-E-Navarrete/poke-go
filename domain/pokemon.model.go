package domain

type Pokemon interface {
	GetID() int
	GetName() string
}

type PokemonEntity struct {
	pokemonID int
	name      string
}

func NewPokemonEntity(pokemonID int, name string) *PokemonEntity {
	return &PokemonEntity{
		pokemonID: pokemonID,
		name:      name,
	}
}

func (p *PokemonEntity) GetID() int {
	return p.pokemonID
}

func (p *PokemonEntity) GetName() string {
	return p.name
}
