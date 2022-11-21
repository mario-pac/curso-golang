package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 //enviando dados para o canal (write)
	<-ch    //recebendo dados do canal (read)
	ch <- 2
	fmt.Println(<-ch)
}
