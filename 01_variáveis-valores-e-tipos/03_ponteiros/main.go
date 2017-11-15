package main

import "fmt"

func main() {
	variável := 10
	fmt.Println(variável)
	// 1. "&" mostra o endereço de uma variável
	fmt.Println(&variável)
	endereço := &variável
	fmt.Println(endereço)
	// 2. "*" faz o de-reference, mostrando o conteúdo da localização pra onde um endereco aponta.
	*endereço++
	fmt.Println(*endereço)
	// 3. E dá pra usar os dois ao mesmo tempo. YOLO.
	fmt.Println(*&variável)
}
