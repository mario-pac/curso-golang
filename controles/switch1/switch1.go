package main

import "fmt"

func notaPraConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default: //quando nenhum dos casos forem saciados
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaPraConceito(9.8))
	fmt.Println(notaPraConceito(6.9))
	fmt.Println(notaPraConceito(2.1))
}
