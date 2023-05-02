package structures

type CreateDeckResponse struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type GetDeckResponse struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

type DrawCardResponse struct {
	Cards []Card `json:"cards"`
}
