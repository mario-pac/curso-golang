package main

import (
	"fmt"
	m "math" //atalho
	_ "time" // usar underline para um import não necessitar de ser utilizado
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // float64 inferido pelo compilador

	//forma reduzida de criar uma var

	area := PI * m.Pow(raio, 2)
	//o go nao deixa que uma variavel seja criada e nao usada
	fmt.Printf("A área da circunferencia é %f.\n", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
