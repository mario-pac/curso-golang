package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		fatorialAnterior := fatorial(n - 1)
		return n * fatorialAnterior
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

	//err := fatorial(-4)
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//}

	//uma solução melhor seria...
}
