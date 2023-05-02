package globals

import (
	"example.com/card-game/structures"
)

//Global variable to store and access decks
var decks = make(map[string]structures.Deck)

//This function returns the deck for the provided id
func GetDeck(deckId string) structures.Deck {
	return decks[deckId]
}

//This function stores/updates the deck for the provided id
func SetDeck(deckId string, deck structures.Deck) {
	decks[deckId] = deck
}
