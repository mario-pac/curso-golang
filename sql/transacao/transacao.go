package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios (id, nome) values ($1, $2)")

	_, err = stmt.Exec(2000, "Bia")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = stmt.Exec(2001, "João")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = stmt.Exec(1, "Mario") //chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
