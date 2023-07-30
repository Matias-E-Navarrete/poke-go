package application

import (
	"github.com/Matias-E-Navarrete/poke-go.git/domain"
)

type GetPokemon struct {
	pokemonApi domain.IPokemonAPI
}

func NewGetPokemon(pokemonApi domain.IPokemonAPI) *GetPokemon {
	return &GetPokemon{pokemonApi: pokemonApi}
}

func (getPokemon *GetPokemon) Execute(pokemonID int) (domain.Pokemon, error) {
	return getPokemon.pokemonApi.GetById(pokemonID)
}
