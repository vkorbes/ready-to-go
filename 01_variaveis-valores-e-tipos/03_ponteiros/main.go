package main

import "fmt"

func main() {
	variável := 10
	fmt.Printf("variável  = %v\n", variável)
	// 1. "&" mostra o endereço de uma variável
	fmt.Printf("&variável = %v\n", &variável)
	endereço := &variável
	fmt.Printf("endereço  = %v\n", endereço)
	// 2. "*" faz o de-reference, mostrando o conteúdo da localização pra onde um endereco aponta.
	*endereço++
	fmt.Printf("variável + 1 = %v\n", *endereço)
	// 3. E dá pra usar os dois ao mesmo tempo. YOLO.
	fmt.Printf("*&variável   = %v\n", *&variável)
}
