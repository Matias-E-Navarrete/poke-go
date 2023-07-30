package pokegogit

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

	http.HandleFunc("/pokemon/", func(w http.ResponseWriter, r *http.Request) {
		pokemonIDStr := r.URL.Query().Get("id")
		pokemonID, error := strconv.Atoi(pokemonIDStr)
		if error != nil {
			http.Error(w, "Invalid Pokémon ID", http.StatusBadRequest)
			return
		}

		data, error := getPokemon.Execute(pokemonID)
		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		// Escribimos los datos del Pokémon en la respuesta HTTP
		_, _ = w.Write([]byte(data.GetName()))

		jsonResponse, error := json.Marshal(data)
		if error != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonResponse)
	})

	fmt.Println("Server Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
