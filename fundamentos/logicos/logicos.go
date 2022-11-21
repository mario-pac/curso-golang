package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // XOR
	tomarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, tomarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Println("Tv50:", tv50, "Tv32:", tv32, "Sorvete:", sorvete, "Saudável:", !sorvete)

}
