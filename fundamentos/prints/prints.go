package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.\n")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor e x é "+ x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %.2f\n", x)

	a, b, c, d := 1, 1.9999, false, "opa"
	fmt.Printf("%d %f %.1f %t %s\n", a, b, b, c, d)
	fmt.Printf("%v %v %v %v\n", a, b, c, d)
}
