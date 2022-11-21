package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("%s falou %d vezes", pessoa, i)
		}
	}()
	return ch
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				ch <- s
			case s := <-entrada2:
				ch <- s
			}
		}
	}()
	return ch
}

func main() {
	c := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
