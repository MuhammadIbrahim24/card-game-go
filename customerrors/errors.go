package customerrors

import (
	"errors"
)

var ErrMalformedCard = errors.New("malformed card")
var ErrMalformedCount = errors.New("malformed count")
var ErrDeckNotFound = errors.New("deck not found")
var ErrInsufficientCards = errors.New("insufficient cards in deck")
var ErrMissingCount = errors.New("missing count")
