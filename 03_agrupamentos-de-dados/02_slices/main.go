package main

import "fmt"

func main() {
	slice := []string{
		"Slices são coleções de valores de um mesmo tipo.",
		"Podemos acessar valores individuais",
	}
	fmt.Print(slice)
	slice = append(slice, "assim.")
	fmt.Println("", slice[2])
	slice = append(slice, "Pra adicionar valores utilizamos a função append.")
	fmt.Println(slice[3])
	outraslice := []string{
		"E dá até pra",
		"adicionar uma slice a outra,",
		"só que aí precisamos desses três pontinhos ali.",
	}
	slice = append(slice, outraslice...)
	fmt.Println("Tudo junto:\n", slice)
	fmt.Println("Tambem dá pra fatiar uma slice:")
	partes := []string{
		"\t...o começo.",
		"\n\t...o meio.",
		"\n\t...o finzinho.",
	}
	tudo := []string{"T", "u", "d", "o!"}
	fmt.Println("Aqui aparece só...\n", partes[:1], partes[1:2], partes[2:])
	fmt.Println("Já aqui, aparece...", tudo[:])
	fmt.Printf("Usando len podemos ver que os slices que criamos até agora tem %d, %d, %d, e %d elementos, respectivamente.\n", len(slice), len(outraslice), len(partes), len(tudo))
	powerrange := []string{
		"Tambem podemos utilizar range",
		"para iterar sobre todos os elementos",
		"de um slice.",
	}
	for i, v := range powerrange {
		fmt.Println(i, v)
	}
	números := []int{80, 111, 114, 32, 101, 110, 113, 117, 97, 110, 116, 111, 32, 117, 116, 105, 108, 105, 122, 97, 109, 111, 115, 32, 115, 111, 109, 101, 110, 116, 101, 32, 115, 116, 114, 105, 110, 103, 115, 44, 32, 109, 97, 115, 32, 112, 111, 100, 101, 109, 111, 115, 32, 116, 101, 114, 32, 115, 108, 105, 99, 101, 115, 32, 100, 101, 32, 116, 111, 100, 111, 115, 32, 111, 115, 32, 116, 105, 112, 111, 115, 46}
	for _, v := range números {
		fmt.Print(string(v))
	}
	inception := [][]string{
		[]string{"Inclusive", "slices", "que"},
		[]string{"contem", "outros", "slices."},
	}
	fmt.Println(inception)
	fmt.Println("Slices", inception[0][1], inception[1][2])
	deletando := []string{
		"Para remover um item",
		"##########",
		"utiliza-se append.",
	}
	deletando = append(deletando[:1], deletando[2:]...)
	fmt.Println(deletando)
}
