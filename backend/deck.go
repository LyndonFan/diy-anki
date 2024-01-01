package backend

import (
	"database/sql"
	"fmt"
	"time"
)

type Deck struct {
	DeckID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func createDeck(db *sql.DB, name string) (int, error) {
	timeNow := time.Now()
	deck := Deck{
		Name:      name,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	createDeckStatement := fmt.Sprintf(
		"INSERT INTO decks (name, created_at, updated_at) VALUES ('%s', '%s', '%s')",
		deck.Name,
		deck.CreatedAt,
		deck.UpdatedAt,
	)
	_, err := db.Exec(createDeckStatement)
	if err != nil {
		return 0, err
	}
	findLastestIDStatement := fmt.Sprintf(
		"SELECT id FROM decks ORDER BY id DESC LIMIT 1",
	)
	res, err := db.Query(findLastestIDStatement)
	if err != nil {
		return 0, err
	}
	// convert res into int
	var deckID int
	for res.Next() {
		err := res.Scan(&deckID)
		if err != nil {
			return 0, err
		}
		return deckID, nil
	}
	return 0, fmt.Errorf("unable to find deck after creating")
}
