package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/card-game/businesslayer"
	"example.com/card-game/customerrors"
	"github.com/gorilla/mux"
)

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	log.Println("Begin: CreateDeck Handler")
	log.Println("Query Parameters:", r.URL.Query())

	var hasCards = r.URL.Query().Has("cards")
	var cards = r.URL.Query().Get("cards")
	var shuffled = r.URL.Query().Get("shuffled")

	var response, err = businesslayer.CreateDeck(hasCards, cards, shuffled)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Response:", response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	log.Println("End: CreateDeck Handler")
}

func GetDeck(w http.ResponseWriter, r *http.Request) {
	log.Println("Begin: GetDeck Handler")

	deckId := mux.Vars(r)["deck_id"]

	var response, err = businesslayer.GetDeck(deckId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Println("Response:", response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	log.Println("End: GetDeck Handler")
}

func DrawCard(w http.ResponseWriter, r *http.Request) {
	log.Println("Begin: DrawCard Handler")

	deckId := mux.Vars(r)["deck_id"]
	count := r.URL.Query().Get("count")

	var response, err = businesslayer.DrawCards(deckId, count)
	if err != nil {
		if err == customerrors.ErrDeckNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	log.Println("Response:", response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	log.Println("End: DrawCard Handler")
}
