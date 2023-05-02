package structures

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type Deck struct {
	Shuffled bool   `json:"shuffled"`
	Cards    []Card `json:"cards"`
}
