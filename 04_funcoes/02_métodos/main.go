package main

import "fmt"

// Métodos são funcionalidade específica para um tipo.

// pizza é nosso método
type pizza struct {
	sabor        string
	preço        float64
	ingredientes string
}

// dados toma pizza como argumento, portanto é um método do tipo argumentos
func (p pizza) dados() {
	fmt.Println("Pizza: ", p.sabor, "\nPreço: R$", p.preço)
}

// receita é outro método do tipo argumentos
func (p pizza) receita() {
	fmt.Println("Mistura", p.ingredientes, "e taca tudo no forno.")
}

func main() {
	abacaxi := pizza{
		sabor:        "Pizza de abacaxi à milanesa",
		preço:        13.37,
		ingredientes: "ovo, abacaxi, palmito, molho de tomate e açafrão",
	}
	pepperoni := pizza{"Pizza de pepperoni sem abacaxi", 21.09, "pepperoni em rodelas, massa fina e molho de tomate especial"}
	// e no nosso programa temos acesso aos métodos do nosso tipo pizza
	abacaxi.dados()
	abacaxi.receita()
	fmt.Println("—")
	pepperoni.dados()
	pepperoni.receita()
}
