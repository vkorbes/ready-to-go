package main

import "fmt"

func main() {
	// 1. Imprime um texto em stdout
	fmt.Print("fmt.Print joga em os.Stdout. ")
	// 2. Imprime um texto com \n implicito no final
	fmt.Println("fmt.Println faz a mesma coisa, mas adiciona uma linha nova.")
	// 3. Copia um texto para uma variável do tipo string
	variável := fmt.Sprintf("fmt.Sprintf")
	// 4. Imprime um texto em stdout substituindo os marcadores % pelos respectivos valores
	fmt.Printf("%v, assim como fmt.Printf, usa format printing. "+
		"Ou seja, toma formatting verbs que podem representar números "+
		"decimais (%d), binários (%b), mostrar tipos (%T), e etc.\n",
		variável, 50, 50, variável)

	// 5. https://golang.org/pkg/fmt/
}
