package main

import "fmt"

// será executado primeiro que tudo
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
