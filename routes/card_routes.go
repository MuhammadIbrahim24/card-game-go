package routes

import (
	"example.com/card-game/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/deck/new", handlers.CreateDeck).Methods("GET")          //Create new deck
	router.HandleFunc("/deck/{deck_id}", handlers.GetDeck).Methods("GET")       //Get deck
	router.HandleFunc("/deck/{deck_id}/draw", handlers.DrawCard).Methods("GET") //Draw card from a given deck
}
