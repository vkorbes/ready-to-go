package main

import "fmt"

// 6. Podemos criar nossos próprios tipos utilizando um tipo subjacente.
type cebola int

var molhodetomate cebola = 9

func main() {
	// 1. Os tipos podem ser "deduzidos" pelo compilador.
	x := 10
	var y = "batata"
	// 2. Ou podem ser declarados explicitamente.
	var z string = "frita"
	var w float64 = 11.1
	// 3. Os tipos de dados primitivos são int, string, bool, e float64.
	var a int
	var b string
	var c bool
	var d float64
	// 4. Os zero-values são:
	// - ints: 0
	// - strings: ""
	// - booleans: false
	// - floats: 0.0
	fmt.Println(a, b, c, d)
	// 5. Os tipos de dados compostos são array, slice, struct, e map.
	var e [5]int
	var f []int
	var g struct{}
	var h map[int]string
	// 7. Para converter valores entre tipos utilizamos tipo(valor).
	x = int(molhodetomate)
	molhodetomate = cebola(x)
	// 8. Em Go não podemos ter variáveis não utilizadas. Portanto pra fazer esse programa rodar precisamos utilizar o que não foi utilizado acima.
	fmt.Println(e, f, g, h, y, z, w)

}
