package main

import "fmt"

func main() {

	// 1. A palavra chave `make` é usada para criar canais
	canal := make(chan string)

	// 2. O operador <- é usado para gravar no canal
	go func() { canal <- "esta mensagem é do canal" }()

	// 3. O operador <-canal é usado para ler do canal
	msg := <-canal
	fmt.Println(msg)
}
