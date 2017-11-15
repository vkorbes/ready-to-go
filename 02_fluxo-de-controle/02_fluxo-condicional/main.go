package main

import "fmt"

func main() {
	for i := 0; i < 13; i++ {
		switch {
		case i == 0:
			fmt.Print("Switch é um condicional ")
		case i == 1:
			fmt.Print("que faz a execução de acordo ")
		case i == 2:
			fmt.Println("com os cases específicados.")
		} // ↓ "i" é o switch statement
		switch i {
		case 3:
			fmt.Print("O switch statement é true por padrão. ")
		case 4:
			fmt.Print("Mas pode ser qualquer coisa, ")
		case 5:
			fmt.Println("e avalia-se se switch statement == case statement.")
		}
		switch j := i + 10; j {
		case 16:
			fmt.Print("Tambem podemos usar uma expressão de ")
		case 17:
			fmt.Println("inicialização antes do switch statement.")
		case 18:
			fmt.Print("Por fim temos o type switch, ")
			type tipo interface{}
			var x tipo = 10
			switch x.(type) {
			case string:
				fmt.Println("me dizendo que `i` tem um valor do tipo string.") // Não...
			case float64:
				fmt.Println("me dizendo que `i` tem um valor do tipo float64.") // Não...
			case int:
				fmt.Println("me dizendo que `i` tem um valor do tipo int.") // Esse roda!
			}
		}
		if i == 9 {
			fmt.Print("\nTambem temos if statements, ")
		} else if i == 10 {
			fmt.Print("que seguem aquela padronagem de if, else if, else. ")
		} else {
			if i == 11 {
				fmt.Println("Dá pra colocar vários uns dentros dos outros, claro.")
			}
		}
		if k := i + 10; k%2 == 0 && k%11 == 0 {
			fmt.Println("E podemos ter uma expressão de inicialização, e podemos encadear mais de uma condição utilizando operadores lógicos condicionais (&&, ||, e família).")
		}
	}
}
