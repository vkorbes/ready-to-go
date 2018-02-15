package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go outra(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}

}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(canal1, canal2 chan int) {
	var wg sync.WaitGroup
	// Pra cada item que canal1 receber, uma nova gofunc é criada que manda resultados pro canal2
	for v := range canal1 {
		wg.Add(1)
		go func(x int) {
			canal2 <- trabalho(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e4)))
	return n
}
