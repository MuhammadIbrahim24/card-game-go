package businesslayer

import (
	"log"
	"strconv"

	"example.com/card-game/customerrors"
	"example.com/card-game/globals"
	"example.com/card-game/structures"
)

func DrawCards(deckId string, count string) (structures.DrawCardResponse, error) {
	log.Println("Begin: DrawCards")
	log.Println("deckId: %s, count:%s", deckId, count)

	//If value of count is not provided return error
	if count == "" {
		return structures.DrawCardResponse{}, customerrors.ErrMissingCount
	}

	//Transform count string to int
	tfCount, err := strconv.Atoi(count)
	if err != nil || tfCount <= 0 {
		return structures.DrawCardResponse{}, customerrors.ErrMalformedCount
	}

	deck := globals.GetDeck(deckId)

	//If Deck doesn't exist in the system
	if globals.GetDeck(deckId).Cards == nil {
		return structures.DrawCardResponse{}, customerrors.ErrDeckNotFound
	}

	//If deck has less cards then the provided count
	if tfCount > len(deck.Cards) {
		return structures.DrawCardResponse{}, customerrors.ErrInsufficientCards
	}

	//Copy drawn cards to new slice and remove from original
	var drawnCards []structures.Card
	drawnCards = append(drawnCards, deck.Cards[:tfCount]...)

	copy(deck.Cards[0:], deck.Cards[tfCount:]) // Shift a[i+1:] left one index.
	deck.Cards = deck.Cards[:len(deck.Cards)-tfCount]
	globals.SetDeck(deckId, deck)

	response := structures.DrawCardResponse{
		Cards: drawnCards,
	}

	log.Println("End: DrawCards")
	return response, nil
}
