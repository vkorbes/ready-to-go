package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type message struct {
	Number int `json:"number"`
}

func main() {
	http.HandleFunc("/", rootHandler)
	// Inicia o servidor escutando na porta 8000
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	// Utiliza a função random para gerar um número aleatório
	random := random()
	// Faz o Marshall do struct
	number, err := json.Marshal(message{random})
	if err != nil {
		log.Fatal(err)
	}
	// Envia o json
	res.WriteHeader(http.StatusOK)
	res.Write(number)
}

func random() int {
	return rand.Intn(1000)
}
