package main

import (
	"fmt"

	"github.com/LyndonFan/diy-anki/backend"
)

// run the backend main function

func main() {
	fmt.Println("Hello, playground from main")
	db, err := backend.GetDataBase()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	res, err := db.Query("SELECT count(*) as \"count\" from decks")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var count int
		err = res.Scan(&count)
		if err != nil {
			panic(err)
		}
		fmt.Println(count)
	}

}
