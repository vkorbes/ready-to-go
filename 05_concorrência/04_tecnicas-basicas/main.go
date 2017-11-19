package main

import "fmt"

// 1. Função que retorna um canal
func geraPares(quantidade int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= quantidade; i++ {
			ch <- i * 2
		}
	}()
	return ch
}

// 3. Função que recebe um canal e devolve outro
func triplicaValores(entrada chan int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for val := range entrada {
			ch <- val * 3
		}
	}()
	return ch
}

func main() {

	// 2. Usando o canal como `generator`
	fmt.Println("Exemplo de generator:")
	meusDados := geraPares(10)
	for val := range meusDados {
		fmt.Printf("Valor: %v\n", val)
	}

	// 4. Exemplo simples de `pipeline`
	fmt.Println("Exemplo de pipeline:")
	meusDados3 := triplicaValores(geraPares(10))
	for val := range meusDados3 {
		fmt.Printf("Valor: %v\n", val)
	}
}
