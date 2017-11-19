package main

import "fmt"

func main() {

	// 1. Declaração de um array vazio
	var tresInteirosVazio [3]int
	fmt.Printf("tresInteirosVazio: %v\n", tresInteirosVazio)

	tresInteirosVazio2 := [3]int{}
	fmt.Printf("tresInteirosVazio2: %v\n", tresInteirosVazio2)

	// 2. Declaração de um array com valores
	tresInteiros := [3]int{1, 2, 3}
	fmt.Printf("tresInteiros: %v\n", tresInteiros)

	// 4. O tamanho também pode ser inferido pelo inicializador:
	quatroFrutas := [...]string{"banana", "laranja", "maçã", "abacate"}
	fmt.Printf("quatroFrutas: %v\n", quatroFrutas)

	// 5. Os operadores len e cap mostram o tamanho atual e a capacidade máxima do array...
	fmt.Printf("len: %v\tcap: %v\ttype: %T\n",
		len(quatroFrutas),
		cap(quatroFrutas),
		quatroFrutas)

	var doces [5]string
	doces[0] = "chocolate"
	doces[1] = "beijinho"
	doces[2] = "brigadeiro"
	fmt.Printf("len: %v\tcap: %v\ttype: %T\n",
		len(doces),
		cap(doces),
		doces)

	// 6. Uma forma fácil de trabalhar com todos os valores de um array é
	// usar um range loop
	for indice, doce := range doces {
		fmt.Println(indice, doce)
	}

	// 7. Se você quiser pode ignorar o indice ou o valor
	for _, doce := range doces {
		fmt.Println(doce)
	}
}
