package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Println(("Servidor rodando em http://localhost:3000"))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
