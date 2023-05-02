package businesslayer

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"example.com/card-game/customerrors"
	"example.com/card-game/globals"
	"example.com/card-game/mappings"
	"example.com/card-game/structures"
	"github.com/google/uuid"
)

// This function contains business logic for deck creation
func CreateDeck(hasCards bool, cards string, shuffled string) (structures.CreateDeckResponse, error) {
	log.Println("Begin: CreateDeck")
	log.Println("hasCards: %s, cards:%s, shuffled:%s\n", hasCards, cards, shuffled)

	var deckCards []structures.Card
	var err error

	//if cards are provided then create partial deck
	if hasCards {
		deckCards, err = createPartialDeck(cards)
		if err != nil {
			return structures.CreateDeckResponse{}, err
		}
	} else {
		deckCards = createFullDeck()
	}

	// Shuffle deck if required
	isShuffled, _ := strconv.ParseBool(shuffled)
	if isShuffled {
		shuffleCards(deckCards)
	}

	//create deck
	deck := structures.Deck{
		Shuffled: isShuffled,
		Cards:    deckCards,
	}

	//generate uuid for deck identification
	deckId := uuid.New().String()

	//store deck in global variable
	globals.SetDeck(deckId, deck)

	//prepare response to return
	response := structures.CreateDeckResponse{
		DeckId:    deckId,
		Shuffled:  isShuffled,
		Remaining: len(deckCards),
	}

	log.Println("End: CreateDeck")
	return response, nil
}

func createPartialDeck(cards string) ([]structures.Card, error) {
	log.Println("Begin: createPartialDeck")
	log.Println("cards:%s\n", cards)

	var splittedCards = strings.Split(cards, ",")
	var deckCards []structures.Card

	//Loop through provided codes to create deck
	for _, code := range splittedCards {
		//When provided card code value not consists of 2 characters
		if len(code) != 2 {
			return []structures.Card{}, customerrors.ErrMalformedCard
		}

		//Get mapping for suits and values
		value := mappings.ValueMapping[code[0:1]]
		suit := mappings.SuitMapping[code[1:2]]
		if value == "" || suit == "" {
			return []structures.Card{}, customerrors.ErrMalformedCard
		}

		//create card object and add in deck
		var card = structures.Card{value, suit, code}
		deckCards = append(deckCards, card)
	}

	log.Println("End: createPartialDeck")
	return deckCards, nil
}

func createFullDeck() []structures.Card {
	log.Println("Begin: createFullDeck")

	var deckCards []structures.Card

	//Loop throuh all suits and values to create complete deck of cards
	for _, suit := range mappings.Suits {
		for _, value := range mappings.Values {
			mappedValue := mappings.ValueMapping[value]
			mappedSuit := mappings.SuitMapping[suit]
			code := value + suit
			var card = structures.Card{mappedValue, mappedSuit, code}
			deckCards = append(deckCards, card)
		}
	}

	log.Println("End: createFullDeck")

	return deckCards
}

func shuffleCards(cards []structures.Card) {
	log.Println("Begin: shuffleCards")
	log.Println("cards:", cards)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	log.Println("End: shuffleCards")
}
