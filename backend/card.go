package backend

import (
	"database/sql"
	"time"
)

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

func CreateCard(db *sql.DB, deckID int, question string, answer string) (int, error) {
	timeNow := time.Now()
	card := Card{
		DeckID:       deckID,
		Question:     question,
		Answer:       answer,
		CreatedAt:    timeNow,
		UpdatedAt:    timeNow,
		LastReviewAt: timeNow,
		NextReviewAt: timeNow,
	}
	createCardStatement := `
		INSERT INTO cards
		(deck_id, question, answer, created_at, updated_at, last_review_at, next_review_at)
		VALUES
		(?, ?, ?, ?, ?, ?, ?);
	`
	execRes, err := db.Exec(
		createCardStatement,
		card.DeckID,
		card.Question,
		card.Answer,
		card.CreatedAt,
		card.UpdatedAt,
		card.LastReviewAt,
		card.NextReviewAt,
	)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := execRes.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastInsertID), nil
}

func DeleteCard(db *sql.DB, cardID int) error {
	_, err := db.Exec("DELETE FROM cards WHERE id = ?;", cardID)
	if err != nil {
		return err
	}
	return nil
}

func FindCards(db *sql.DB, deckID int) ([]Card, error) {
	rows, err := db.Query("SELECT id, question, answer FROM cards WHERE deck_id = ?;", deckID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cards []Card
	for rows.Next() {
		var card Card
		err = rows.Scan(&card.CardID, &card.Question, &card.Answer)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
