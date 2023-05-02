package testcases

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"example.com/card-game/globals"
	"example.com/card-game/structures"
)

func TestDrawCardFromNonExistentDeck(t *testing.T) {
	fmt.Println("**Draw-Card Test 1: Draw Card From Non-Existent Deck**")
	req, _ := http.NewRequest("GET", URI+"/d2604cab-b658-43b2-ae35-3e0472d9c42a/draw?count=3", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	if body := response.Body.String(); body != "deck not found\n" {
		t.Errorf("Expected error 'deck not found'. Got %s", response.Body.String())
	}
}

func TestDrawCardWhenMissingCount(t *testing.T) {
	fmt.Println("**Draw-card Test 2: Draw Card When count parameter is missing**")
	req, _ := http.NewRequest("GET", URI+"/d2604cab-b658-43b2-ae35-3e0472d9c42a/draw", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if body := response.Body.String(); body != "missing count\n" {
		t.Errorf("Expected error 'missing count'. Got %s", response.Body.String())
	}
}

func TestDrawCardsMoreThanInDeck(t *testing.T) {
	fmt.Println("**Draw-card Test 3: Draw Cards more than the count present in deck**")

	//creating a new deck
	req, _ := http.NewRequest("GET", URI+"/new?cards=AS,KD,2H,JS", nil)
	response := executeRequest(req)
	result := structures.CreateDeckResponse{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	req, _ = http.NewRequest("GET", URI+"/"+result.DeckId+"/draw?count=5", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if body := response.Body.String(); body != "insufficient cards in deck\n" {
		t.Errorf("Expected error 'insufficient cards in deck'. Got %s", response.Body.String())
	}
}

func TestDrawCardsFromDeck(t *testing.T) {
	fmt.Println("**Draw-card Test 4: Draw Cards from Deck**")

	//creating a new deck
	req, _ := http.NewRequest("GET", URI+"/new?cards=AS,KD,2H,JS", nil)
	response := executeRequest(req)
	result := structures.CreateDeckResponse{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	req, _ = http.NewRequest("GET", URI+"/"+result.DeckId+"/draw?count=2", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	drawCardResult := structures.DrawCardResponse{}
	if err := json.NewDecoder(response.Body).Decode(&drawCardResult); err != nil {
		log.Fatalln(err)
	}

	if len(drawCardResult.Cards) != 2 {
		t.Errorf("Expected cards array to be of length 2. Got %d", len(drawCardResult.Cards))
	}

	deck := globals.GetDeck(result.DeckId)

	if len(deck.Cards) != 2 {
		t.Errorf("Expected deck length to be reduced to 2 after withdrawal. Got %d", len(deck.Cards))
	}
}
