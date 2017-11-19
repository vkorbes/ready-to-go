package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 1. Invocação normal de função
	f("chamada direta")

	// 2. Para criar uma goroutine e executar em paralelo basta colocar a palavra `go` na frente
	go f("goroutine")

	// 3. Pode ser uma função anônima
	go func(msg string) {
		fmt.Println(msg)
	}("goroutine com função anônima")

	// 4. Termina quando o usuário apertar `enter`
	var input string
	fmt.Scanln(&input)
	fmt.Println("fim!")
}
