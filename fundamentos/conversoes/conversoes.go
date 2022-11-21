package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFInal := int(nota)
	fmt.Println(notaFInal)

	//cuidado...
	fmt.Println("Teste " + string(123))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(97))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 222)

	b, _ := strconv.ParseBool("1") //ou true
	if b {
		fmt.Println("Verdadeiro")
	}
}
