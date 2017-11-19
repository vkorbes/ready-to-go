package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := soma(20, 30)
	if resultado != 50 {
		t.Errorf("esperado: %v recebido: %v\n", 50, resultado)
	}
}

func TestSomaTabela(t *testing.T) {
	tabela := []struct {
		entrada1 int
		entrada2 int
		esperado int
	}{
		{
			1,
			2,
			3,
		},
		{
			2,
			4,
			6,
		},
		{
			-2,
			-4,
			-5,
		},
	}

	for numero, teste := range tabela {
		resultado := soma(teste.entrada1, teste.entrada2)
		if teste.esperado != resultado {
			t.Errorf("teste: %v esperado: %v recebido: %v\n", numero, teste.esperado, resultado)
		}
	}
}
