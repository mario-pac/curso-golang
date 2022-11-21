package main

import "fmt"

func main() {
	i := 1

	//Go nao tem aritmetica de ponteiros
	var p *int = nil
	p = &i //pegando endereço da variável
	*p++   //desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
