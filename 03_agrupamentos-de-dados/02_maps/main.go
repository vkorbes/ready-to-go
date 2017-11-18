package main

import "fmt"

func main() {
	m := map[string]string{
		"Maps são pares": "de chaves e valores.",
	}
	fmt.Println(m)
	n := map[int]string{
		2: "A gente acessa",
		3: "valores em maps",
		5: "assim...",
	}
	fmt.Println(n[2], n[3], n[5])
	o := map[string]int{
		"valorX": 2,
		"valorY": 1,
		"valorZ": 0,
	}
	fmt.Println("Estes valores existem:", o["valorX"], o["valorY"], o["valorZ"])
	fmt.Println("Mas e este?", o["valorW"])
	fmt.Println("Como eu faço pra saber se o key existe e o valor é zero, ou se ele não existe?")
	fmt.Print("Assim: ")
	z, ok := o["valorZ"]
	fmt.Println("O valorZ existe?", ok, "E seu valor é:", z)
	w, ok := o["valorW"]
	fmt.Println("O valorW existe?", ok, "Portanto, ele retorna:", w)
	fmt.Println("Na prática, utilizamos assim:")
	p := map[string]string{"valor": "existe"}
	if v, ok := p["valor"]; ok {
		fmt.Println("O valor", v, "(", ok, ")")
	}
	q := map[int]string{
		1:   "Para apagar",
		2:   "um valor",
		3:   "utilizamos delete().",
		666: "IRON MAIDEN MANO!!!!!11",
	}
	delete(q, 666)
	fmt.Println(q[1], q[2], q[3])
	fmt.Print("E podemos utilizar range com maps, mas a ordem... ")
	r := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	for _, x := range r {
		fmt.Print(x, " ")
	}
	fmt.Println("...é aleatória.")
}
