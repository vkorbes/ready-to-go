package main

import "fmt"

func main() {
	// 1. Canal `unbuffered`
	ch1 := make(chan int)
	// 2. Canal `buffered`
	ch2 := make(chan int, 3)

	go func() {
		defer close(ch1)
		// 3. Se você comentar esta linha...
		ch1 <- 1
	}()

	// 4. ... o código vai travar aqui -> deadlock
	val, ok := <-ch1
	fmt.Printf("Lendo de ch1 (aberto): %v, %v\n", val, ok)

	go func() {
		// 5. Esquecer de fechar um canal também pode causar deadlock
		defer close(ch2)
		ch2 <- 2
		ch2 <- 3
		ch2 <- 4
		ch2 <- 5
	}()

	// 6. Se você tentar gravar no canal bufered mais valores do que a sua capacidade
	// o canal também vai bloquear (experimente)

	for val := range ch2 {
		fmt.Printf("ch2: %v\n", val)
	}

	// 7. Ler de um canal fechado
	val, ok = <-ch1
	fmt.Printf("Lendo de ch1 (fechado): %v, %v\n", val, ok)
}
