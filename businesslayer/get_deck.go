package businesslayer

import (
	"log"

	"example.com/card-game/customerrors"
	"example.com/card-game/globals"
	"example.com/card-game/structures"
)

func GetDeck(deckId string) (structures.GetDeckResponse, error) {
	log.Println("Begin: GetDeck")
	log.Println("deckId: %s", deckId)

	//If deck doesn't exist in system
	deck := globals.GetDeck(deckId)
	if deck.Cards == nil {
		return structures.GetDeckResponse{}, customerrors.ErrDeckNotFound
	}

	response := structures.GetDeckResponse{
		DeckId:    deckId,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
		Cards:     deck.Cards,
	}

	log.Println("End: GetDeck")
	return response, nil
}
