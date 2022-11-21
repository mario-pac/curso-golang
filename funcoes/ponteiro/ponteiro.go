package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	//acessando o valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) //por valor
	fmt.Println(n)

	//& para endereço de variável
	inc2(&n) //por referencia
	fmt.Println(n)
}
