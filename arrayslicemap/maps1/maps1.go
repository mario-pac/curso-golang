package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//mapas devem ser inicializados

	aprovados := make(map[int]string)
	aprovados[12345678] = "Maria"
	aprovados[87654321] = "Pedro"
	aprovados[19283746] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Println(aprovados[19283746])
	delete(aprovados, 19283746)
	fmt.Println("oi", aprovados[19283746])

}
