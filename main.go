package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"dojo.com/rest-api/middlewares"
	"github.com/gorilla/mux"
)

type JokenpoResponse struct {
	Winner      string `json:"winner"`
	Description string `json:"description"`
}

var result JokenpoResponse
var options map[string]int

func validateParameters(player1 string, player2 string) (string, string, error) {
	options = map[string]int{
		"PEDRA":   1,
		"PAPEL":   2,
		"TESOURA": 3,
	}

	player1 = strings.ToUpper(player1)
	player2 = strings.ToUpper(player2)

	if options[player1] == 0 || options[player2] == 0 {
		return "", "", errors.New("os valores estão incorretos")
	}

	return player1, player2, nil
}

func Jokenpo(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	p1, p2, err := validateParameters(v.Get("player1"), v.Get("player2"))

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result = JokenpoResponse{}

	switch {
	case options[p1] == 1 && options[p2] == 3:
		result = JokenpoResponse{Winner: "player1", Description: fmt.Sprintf("%s ganha de %s", p1, p2)}
	case options[p1] == 3 && options[p2] == 1:
		result = JokenpoResponse{Winner: "player2", Description: fmt.Sprintf("%s ganha de %s", p2, p1)}
	case options[p1] > options[p2]:
		result = JokenpoResponse{Winner: "player1", Description: fmt.Sprintf("%s ganha de %s", p1, p2)}
	case options[p1] < options[p2]:
		result = JokenpoResponse{Winner: "player2", Description: fmt.Sprintf("%s ganha de %s", p2, p1)}
	default:
		result = JokenpoResponse{Description: fmt.Sprintf("%s empata com %s", p1, p2)}
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.ValidateMiddleware)
	router.HandleFunc("/jokenpo", Jokenpo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
