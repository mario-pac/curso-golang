package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//exec(db, "create database cursogo")
	//exec(db, `create table usuarios(
	//	id integer,
	//	nome varchar(80)
	//)`)

	// CRIANDO INSERTS
	stmt, _ := db.Prepare("insert into usuarios(id, nome) values($1, $2)")
	stmt.Exec(0, "Maria")
	stmt.Exec(1, "João")

	res, _ := stmt.Exec(2, "Pedro")

	id, _ := res.LastInsertId() //nao funciona no postgres
	fmt.Println("Ultimo id inserido:", id)

	linhas, _ := res.RowsAffected()
	fmt.Println("Linhas afetadas: ", linhas)
}
