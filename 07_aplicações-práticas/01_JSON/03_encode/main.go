package main

import "os"
import "encoding/json"

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {
	jamesbond := pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 1000000.50,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesbond)

}
