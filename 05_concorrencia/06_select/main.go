package main

import (
	"fmt"
	"time"
)

func umProcessoPesado() chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- "demorou mas deu certo =)"
	}()
	return ch
}

// 1. Rodando um processo "sequencial"
func main() {
	ch := umProcessoPesado()
	fmt.Println(<-ch)
}

// 2. Mesma coisa usando o select = blocking
// func main() {
// 	ch := umProcessoPesado()

// 	select {
// 	case val := <-ch:
// 		fmt.Println(val)
// 	}
// }

// 3. Com valor default = non-blocking
// func main() {
// 	ch := umProcessoPesado()

// 	select {
// 	case val := <-ch:
// 		fmt.Println(val)
// 	default:
// 		//fmt.Println("Não tenho paciência")
// 	}
// }

// 4. Com timeout
// func main() {
// 	ch := umProcessoPesado()

// 	select {
// 	case val := <-ch:
// 		fmt.Println(val)
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("Acabou o tempo")
// 	}
// }

// 5. Selecionando entre vários canis
// func main() {
// 	ch := umProcessoPesado()
// 	ch2 := umProcessoPesado()

// 	select {
// 	case val := <-ch:
// 		fmt.Println(val)
// 	case val2 := <-ch2:
// 		fmt.Println("outro val", val2)
// 	}
// }
