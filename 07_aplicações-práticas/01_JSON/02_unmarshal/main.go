package main

import (
	"encoding/json"
	"fmt"
)

type informacoes struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

func main() {

	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":1000000.5}`)

	var jamesbond informacoes
	err := json.Unmarshal(sb, &jamesbond)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(jamesbond)
	fmt.Println(jamesbond.Profissao)
}
