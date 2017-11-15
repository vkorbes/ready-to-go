package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Aqui temos um for loop.")
	}
	i := 0
	for i < 3 {
		fmt.Println("Aqui temos um for loop. (Que é como Go chama um while.)")
		i++
	}
	for {
		fmt.Println("E esse é um loop infinito. Ou seria, se não houvesse um `break`.")
		i++
		if i > 5 {
			break
		}
	}
	for {
		fmt.Println("Com `continue` podemos ir direto para a próxima iteração.")
		i++
		if i < 9 {
			continue
		}
		fmt.Println("Por isso que isso aqui aparece uma vez só.")
		return
	}
}
