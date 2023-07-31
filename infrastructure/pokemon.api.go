package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Matias-E-Navarrete/poke-go.git/domain"
)

const baseURL = "https://pokeapi.co/api/v2/"

type PokemonAPI struct{}

func NewPokeAPI() *PokemonAPI {
	return &PokemonAPI{}
}

func (api *PokemonAPI) GetById(pokemonId int) (domain.Pokemon, error) {
	apiURL := fmt.Sprintf("%s/pokemon/%d", baseURL, pokemonId)
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("we have a problem with api connnection")
	}

	var dataResponse struct {
		PokemonID int    `json:"id"`
		Name      string `json:"name"`
	}

	if error := json.NewDecoder(response.Body).Decode(&dataResponse); error != nil {
		return nil, error
	}

	var pokemon = domain.NewPokemonEntity(dataResponse.PokemonID, dataResponse.Name)
	return pokemon, nil
}
