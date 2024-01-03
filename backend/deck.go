package backend

import (
	"database/sql"
	"time"
)

type Deck struct {
	DeckID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateDeck(db *sql.DB, name string) (int, error) {
	timeNow := time.Now()
	deck := Deck{
		Name:      name,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	createDeckStatement := "INSERT INTO decks (name, created_at, updated_at) VALUES (?, ?, ?);"
	execRes, err := db.Exec(createDeckStatement, deck.Name, deck.CreatedAt, deck.UpdatedAt)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := execRes.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastInsertID), nil
}

func UpdateDeck(db *sql.DB, deckID int, name string) error {
	timeNow := time.Now()
	updateDeckStatement := "UPDATE decks SET name = ?, updated_at = ? WHERE id = ?;"
	_, err := db.Exec(updateDeckStatement, name, timeNow, deckID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDeck(db *sql.DB, deckID int) error {
	_, err := db.Exec("DELETE FROM cards WHERE deck_id = ?;", deckID)
	if err != nil {
		return err
	}
	return nil
}
