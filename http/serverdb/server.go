package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("Executando na url http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
