package main

import (
	"fmt"
)

// Elevador é uma coisa que sobe e desce
type Elevador interface {
	Sobe()
	Desce()
}

// Escada é um jeito bom de gastar calorias
type Escada struct {
	degraus  int
	calorias int
}

// Sobe a escada
func (e Escada) Sobe() {
	fmt.Printf("Subi %v degraus e gastei %v calorias\n", e.degraus, e.calorias)
}

// Desce a escada
func (e Escada) Desce() {
	fmt.Printf("Desci %v degraus e gastei %v calorias\n", e.degraus, e.calorias/2)
}

// SobeAndar é usa a interface como parâmetro
func SobeAndar(elev Elevador) {
	elev.Sobe()
}

// DesceAndar é usa a interface como parâmetro
func DesceAndar(elev Elevador) {
	elev.Desce()
}

// ElevadorPanoramico é mais rapido do que escada (e o nome Elevador já estava sendo usado XD)
type ElevadorPanoramico struct {
	velocidade int
}

// Sobe de elevador panoramico
func (ep ElevadorPanoramico) Sobe() {
	fmt.Println("Subi de andar apreciando a vista, mas não emagreci nada")
}

// Desce de elevador panoramico
func (ep ElevadorPanoramico) Desce() {
	fmt.Println("Desci de andar apreciando a vista, mas não emagreci nada")
}

// Implementa a stringer interface
func (ep ElevadorPanoramico) String() string {
	return fmt.Sprintf("[velocidade: %v]", ep.velocidade)
}

func main() {
	// 1. Escada implementa a interface Elevador
	e := Escada{100, 10}
	SobeAndar(e)
	DesceAndar(e)

	// 2. E ElevadorPanoramico também!
	ep := ElevadorPanoramico{20}
	SobeAndar(ep)
	DesceAndar(ep)

	// 3. ElevadorPanoramico implementa a interface "Stringer"
	fmt.Printf("%v\n", ep)
}
