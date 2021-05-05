package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type JokenpoResponse struct {
	Winner      string `json:"winner"`
	Description string `json:"description"`
}

var result JokenpoResponse

func Jokenpo(w http.ResponseWriter, r *http.Request) {
	result = JokenpoResponse{Winner: "player1", Description: "pedra ganha de tesoura"}
	json.NewEncoder(w).Encode(result)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/jokenpo", Jokenpo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
