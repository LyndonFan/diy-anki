package backend

import "time"

type Card struct {
	CardID       int
	DeckID       int
	Question     string
	Answer       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastReviewAt time.Time
	NextReviewAt time.Time
}
