package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Matias-E-Navarrete/poke-go.git/application"
	"github.com/Matias-E-Navarrete/poke-go.git/infrastructure"
)

func main() {
	// Create instance HTTP With PokeAPI
	httpAdapter := infrastructure.NewPokeAPI()

	// Create instance use case.
	getPokemon := application.NewGetPokemon(httpAdapter)

	http.HandleFunc("/pokemon/", func(response http.ResponseWriter, request *http.Request) {
		pokemonIDStr := request.URL.Path[len("/pokemon/"):]
		pokemonID, error := strconv.Atoi(pokemonIDStr)
		if error != nil {
			http.Error(response, "Invalid Pokémon ID", http.StatusBadRequest)
			return
		}

		data, error := getPokemon.Execute(pokemonID)
		if error != nil {
			http.Error(response, error.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		pokemonJSON := map[string]interface{}{
			"id":   data.GetID(),
			"name": data.GetName(),
		}
		jsonResponse, error := json.Marshal(pokemonJSON)
		if error != nil {
			http.Error(response, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
		response.Header().Set("Content-Type", "application/json")
		_, _ = response.Write(jsonResponse)
	})

	fmt.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
