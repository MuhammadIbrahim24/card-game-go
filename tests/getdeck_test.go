package testcases

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"example.com/card-game/structures"
)

func TestGetNonExistingDeck(t *testing.T) {
	fmt.Println("**Get-Deck Test 1: Get Non-Existing Deck**")
	req, _ := http.NewRequest("GET", URI+"/d2604cab-b658-43b2-ae35-3e0472d9c42a", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	if body := response.Body.String(); body != "deck not found\n" {
		t.Errorf("Expected error 'deck not found'. Got %s", response.Body.String())
	}
}

func TestGetUnpassedDeck(t *testing.T) {
	fmt.Println("**Get-Deck Test 2: Get Unpassed Deck**")
	req, _ := http.NewRequest("GET", URI, nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	if body := response.Body.String(); body != "404 page not found\n" {
		t.Errorf("Expected error '404 page not found'. Got %s", response.Body.String())
	}
}

func TestGetExistingDeck(t *testing.T) {
	fmt.Println("**Get-Deck Test 3: Get Existing Deck**")

	//creating a new deck
	req, _ := http.NewRequest("GET", URI+"/new", nil)
	response := executeRequest(req)
	result := structures.CreateDeckResponse{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	//getting created deck
	req, _ = http.NewRequest("GET", URI+"/"+result.DeckId, nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	getDeckResult := structures.GetDeckResponse{}
	if err := json.NewDecoder(response.Body).Decode(&getDeckResult); err != nil {
		log.Fatalln(err)
	}

	if getDeckResult.Remaining != 52 {
		t.Errorf("Expected remaining cards to be 52. Got %d", getDeckResult.Remaining)
	}
	if getDeckResult.Shuffled != false {
		t.Errorf("Expected Shuffled to be false. Got %t", getDeckResult.Shuffled)
	}
	if getDeckResult.DeckId == "" {
		t.Errorf("Expected Deck Id to be non-empty. Got empty")
	}
	if len(getDeckResult.Cards) != 52 {
		t.Errorf("Expected cards array to be of length 52. Got %d", len(getDeckResult.Cards))
	}
}
