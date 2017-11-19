package main

import "fmt"

// 2. A keyword var cria novas variáveis. Pode-se usar onde quiser. Fora de code blocks cria variáveis com package level scope.
var abóbora string = "que fome"

func main() {
	// 1. O gopher operator cria novas variáveis. Só pode ser usado dentro de code block. O scope é local.
	abobrinha := "que fome"
	comida()
	// 3. O operador de atribuição dá um novo valor a uma variável já existente. Não é a mesma coisa que o gopher.
	abobrinha = "que fome!!!!1"
	fmt.Println(abobrinha)
}

func comida() {
	// 4. Isso funciona pois abóbora tem package level scope.
	fmt.Println(abóbora)
	// 5. Isso não roda pois o scope de abobrinha é somente a função main.
	// fmt.Println(abobrinha)
}
