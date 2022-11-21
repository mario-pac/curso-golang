package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, it := range p.itens {
		total += it.preco * float64(it.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{3, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$%.2f\n", pedido.valorTotal())

}
