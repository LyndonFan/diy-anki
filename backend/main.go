package backend

import "fmt"

func main() {
	db, err := GetDataBase()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Hello, playground")
	res, err := db.Query("SELECT * FROM decks")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var id int
		var name string
		var createdAt string
		var updatedAt string
		err = res.Scan(&id, &name, &createdAt, &updatedAt)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name, createdAt, updatedAt)
	}
}
