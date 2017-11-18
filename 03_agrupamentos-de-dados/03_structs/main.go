package main

import "fmt"

type éassimque struct {
	sedeclara int
	umstruct  bool
}

type starwars struct {
	ano  int
	nome string
}

func main() {
	esperança := starwars{
		ano:  1977,
		nome: "Star Wars: A New Hope",
	}
	// esperança := starwars{1977, "Star Wars: A New Hope"} ← mesma coisa!
	fmt.Println("O melhor filme de", esperança.ano, "foi", esperança.nome)
	type meustructbonitinho struct {
		a string
		b string
	}
	type frases struct {
		frase1 []string
		frase2 []string
		frase3 map[int]string
		frase4 meustructbonitinho
	}
	voutecontar := frases{
		frase1: []string{"Podemos", "inclusive", "ter"},
		frase2: []string{"slices", "dentro", "de", "structs."},
		frase3: map[int]string{
			1: "E maps",
			2: "(aleatórios!)",
			3: "tambem",
			4: "pode.",
		},
		frase4: meustructbonitinho{
			a: "E olha só...",
			b: "até structs! (Inception!)",
		},
	}
	fmt.Println(voutecontar)
	anônimo := struct {
		porfim     string
		parenteses string
	}{
		porfim:     "Por fim, temos os structs anônimos",
		parenteses: "(que são, tipo, descartáveis).",
	}
	fmt.Println(anônimo)
}
