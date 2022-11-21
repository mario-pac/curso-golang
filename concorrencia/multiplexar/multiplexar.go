package main

import (
	"cursogolang/pacote/html"
	"fmt"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://wwww.cod3r.com.br", "www.google.com"),
		html.Titulo("https://wwww.amazon.com.br", "https://www.youtube.com"),
	)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
