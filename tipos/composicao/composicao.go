package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoluxuoso interface {
	esportivo
	luxuoso
	//pode adicionar mais métodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("TURBO")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("BALIZA")
}

func main() {
	var b esportivoluxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
