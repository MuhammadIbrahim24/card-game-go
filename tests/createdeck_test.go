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

func TestDefaultUnshuffledFullDeckCreation(t *testing.T) {
	fmt.Println("**Create-Deck Test 1: Create default Unshuffled Full Deck**")
	req, _ := http.NewRequest("GET", URI+"/new", nil)
	response := executeRequest(req)

	result := structures.CreateDeckResponse{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	checkResponseCode(t, http.StatusOK, response.Code)

	if result.Remaining != 52 {
		t.Errorf("Expected remaining cards to be 52. Got %d", result.Remaining)
	}
	if result.Shuffled != false {
		t.Errorf("Expected Shuffled to be false. Got %t", result.Shuffled)
	}
	if result.DeckId == "" {
		t.Errorf("Expected Deck Id to be non-empty. Got empty")
	}

	//Check first 2 cards in deck to verify that the deck is unshuffled
	deck := globals.GetDeck(result.DeckId)
	if deck.Cards[0].Code != "AS" || deck.Cards[1].Code != "2S" {
		t.Errorf("Created deck is not unshuffled")
	}
}

func TestUnshuffledFullDeckCreation(t *testing.T) {
	fmt.Println("**Create-Deck Test 2: Create Unshuffled Full Deck**")
	req, _ := http.NewRequest("GET", URI+"/new?shuffled=false", nil)
	response := executeRequest(req)

	result := structures.CreateDeckResponse{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	checkResponseCode(t, http.StatusOK, response.Code)

	if result.Remaining != 52 {
		t.Errorf("Expected remaining cards to be 52. Got %d", result.Remaining)
	}
	if result.Shuffled != false {
		t.Errorf("Expected Shuffled to be false. Got %t", result.Shuffled)
	}
	if result.DeckId == "" {
		t.Errorf("Expected Deck Id to be non-empty. Got empty")
	}

	//Check first 2 cards in deck to verify that the deck is unshuffled
	deck := globals.GetDeck(result.DeckId)
	if deck.Cards[0].Code != "AS" || deck.Cards[1].Code != "2S" {
		t.Errorf("Created deck is not unshuffled")
	}
}

func TestShuffledFullDeckCreation(t *testing.T) {
	fmt.Println("**Create-Deck Test 3: Create Shuffled Full Deck**")
	req, _ := http.NewRequest("GET", URI+"/new?shuffled=true", nil)
	response := executeRequest(req)

	result := structures.CreateDeckResponse{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	checkResponseCode(t, http.StatusOK, response.Code)

	if result.Remaining != 52 {
		t.Errorf("Expected remaining cards to be 52. Got %d", result.Remaining)
	}
	if result.Shuffled != true {
		t.Errorf("Expected Shuffled to be true. Got %t", result.Shuffled)
	}
	if result.DeckId == "" {
		t.Errorf("Expected Deck Id to be non-empty. Got empty")
	}

	//Check first 2 cards in deck to verify that the deck is shuffled
	deck := globals.GetDeck(result.DeckId)
	if deck.Cards[0].Code == "AS" && deck.Cards[1].Code == "2S" {
		t.Errorf("Created deck is not shuffled")
	}
}

func TestUnshuffledPartialDeckCreation(t *testing.T) {
	fmt.Println("**Create-Deck Test 4: Create Unshuffled Partial Deck**")
	req, _ := http.NewRequest("GET", URI+"/new?shuffled=false&cards=AS,KD,AC,2C,KH", nil)
	response := executeRequest(req)

	result := structures.CreateDeckResponse{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	checkResponseCode(t, http.StatusOK, response.Code)

	if result.Remaining != 5 {
		t.Errorf("Expected remaining cards to be 52. Got %d", result.Remaining)
	}
	if result.Shuffled != false {
		t.Errorf("Expected Shuffled to be false. Got %t", result.Shuffled)
	}
	if result.DeckId == "" {
		t.Errorf("Expected Deck Id to be non-empty. Got empty")
	}

	//Check first 2 cards in deck to verify that the deck is shuffled
	deck := globals.GetDeck(result.DeckId)
	if deck.Cards[0].Code != "AS" || deck.Cards[1].Code != "KD" {
		t.Errorf("Created deck is not unshuffled")
	}
}

func TestShuffledPartialDeckCreation(t *testing.T) {
	fmt.Println("**Create-Deck Test 5: Create Shuffled Partial Deck**")
	req, _ := http.NewRequest("GET", URI+"/new?shuffled=true&cards=AS,KD,AC,2C,KH", nil)
	response := executeRequest(req)

	result := structures.CreateDeckResponse{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	checkResponseCode(t, http.StatusOK, response.Code)

	if result.Remaining != 5 {
		t.Errorf("Expected remaining cards to be 52. Got %d", result.Remaining)
	}
	if result.Shuffled != true {
		t.Errorf("Expected Shuffled to be false. Got %t", result.Shuffled)
	}
	if result.DeckId == "" {
		t.Errorf("Expected Deck Id to be non-empty. Got empty")
	}

	//Check first 2 cards in deck to verify that the deck is shuffled
	deck := globals.GetDeck(result.DeckId)
	if deck.Cards[0].Code == "AS" && deck.Cards[1].Code == "KD" {
		t.Errorf("Created deck is not shuffled")
	}
}

func TestMalformedCardValue(t *testing.T) {
	fmt.Println("**Create-Deck Test 6: Malformed Card Value**")
	req, _ := http.NewRequest("GET", URI+"/new?shuffled=true&cards=AF,KD,AC,2C,KH", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
	if body := response.Body.String(); body != "malformed card\n" {
		t.Errorf("Expected error 'malformed code'. Got %s", response.Body.String())
	}
}
