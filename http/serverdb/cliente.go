package main

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

// Usuario
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para a função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > -1:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... não foi possível terminar seu processo")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //return 500
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("select * from usuarios where id = $1", id).Scan(&u.ID, &u.Nome) //parametros que serao retornados do usuario no scan

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //return 500
	}
	defer db.Close()

	rows, _ := db.Query("select * from usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}
	json, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
