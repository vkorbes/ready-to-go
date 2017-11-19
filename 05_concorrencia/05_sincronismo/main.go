package main

import (
	"fmt"
	"time"
)

// 1. Programando sequencialmente
func fazAlgoImportante(s string) {
	time.Sleep(1 * time.Second)
	fmt.Println(s)
}

func fazTudo(tarefas []string) {
	for _, tarefa := range tarefas {
		fazAlgoImportante(tarefa)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	tarefas := []string{"dormir", "comer", "correr", "trabalhar", "comer", "dormir"}
	fazTudo(tarefas)
}

// 2. Programando em paralelo utilizando um canal para sincronizar
// func fazTudo(tarefas []string, done chan bool) {
// 	for _, tarefa := range tarefas {
// 		fazAlgoImportante(tarefa)
// 		time.Sleep(10 * time.Millisecond)
// 	}
// 	done <- true
// }

// func main() {
// 	tarefas := []string{"dormir", "comer", "correr", "trabalhar", "comer", "dormir"}
// 	done := make(chan bool)
// 	go fazTudo(tarefas, done)
// 	<-done
// }

// 3. Programando em paralelo utilizando um waitgroups para sincronizar
// func fazTudo(tarefas []string) {
// 	var wg sync.WaitGroup
// 	for _, tarefa := range tarefas {
// 		wg.Add(1)
// 		go func(t string) {
// 			defer wg.Done()
// 			fmt.Println(t)
// 		}(tarefa)
// 		time.Sleep(10 * time.Millisecond)
// 	}
// 	wg.Wait()
// }

// func main() {
// 	tarefas := []string{"dormir", "comer", "correr", "trabalhar", "comer", "dormir"}
// 	fazTudo(tarefas)
// }
