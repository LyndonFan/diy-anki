package main

import (
	"database/sql"
	"fmt"

	"github.com/LyndonFan/diy-anki/backend"
)

// run the backend main function

func countDecks(db *sql.DB) (int, error) {
	res, err := db.Query("SELECT count(*) as \"count\" from decks;")
	if err != nil {
		return 0, err
	}
	defer res.Close()
	var count int
	for res.Next() {
		err = res.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}

func countCards(db *sql.DB, deckID int) (int, error) {
	query := fmt.Sprintf(`
		SELECT count(*) as "count"
		FROM cards
		WHERE deck_id = %d;
	`,
		deckID,
	)
	res, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	defer res.Close()
	var count int
	for res.Next() {
		err = res.Scan(&count)
		if err != nil {
			return 0, err
		}
	}
	return count, nil
}

func main() {
	fmt.Println("Hello, playground from main")
	db, err := backend.GetDataBase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	numDecks, err := countDecks(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of decks:", numDecks)

	deckID, err := backend.CreateDeck(db, "test")
	if err != nil {
		panic(err)
	}
	_, err = backend.CreateCard(db, deckID, "foo0", "bar0")
	if err != nil {
		panic(err)
	}
	_, err = backend.CreateCard(db, deckID, "foo1", "bar1")
	if err != nil {
		panic(err)
	}

	numDecks, err = countDecks(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of decks:", numDecks)

	numCards, err := countCards(db, deckID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of cards:", numCards)
}
