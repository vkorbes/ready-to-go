package main

import "fmt"

func main() {
	fmt.Print("fmt.Print joga em os.Stdout. ")
	fmt.Println("fmt.Println faz a mesma coisa, mas adiciona uma linha nova.")
	variável := fmt.Sprintf("fmt.Sprintf")
	fmt.Printf("%v, assim como fmt.Printf, usa format printing. Ou seja, toma formatting verbs que podem representar números decimais (%d), binários (%b), mostrar tipos (%T), e etc.\n", variável, 50, 50, variável)
}
