package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > $1", 0)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)

	}
}
