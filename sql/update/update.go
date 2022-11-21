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

	//update
	stmt, _ := db.Prepare("update usuarios set nome=$1 where id=$2")
	stmt.Exec("Uóxinton Istive", 2)
	stmt.Exec("Sheristone Ualeska", 3)

	//delete
	stmt2, _ := db.Prepare("delete from usuarios where id=$1")
	stmt2.Exec("2000")
}
