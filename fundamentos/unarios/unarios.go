package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++ // x+=1
	y-- // y-=1

	fmt.Println(x, y)

	//fmt.Println(++x == y--) isso nao pode, somente postfix e sem estar em expressões
}
