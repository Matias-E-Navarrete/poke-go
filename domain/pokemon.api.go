package domain

type IPokemonAPI interface {
	GetById(pokemonId int) (Pokemon, error)
}
