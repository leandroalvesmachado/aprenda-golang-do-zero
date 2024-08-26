package main

import (
	"fmt"
	"math"
)

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	raio float64
}

type Forma interface {
	area() float64
}

// obrigatorio ter esse metodo com a mesma assinatura da interface
func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

// obrigatorio ter esse metodo com a mesma assinatura da interface
func (c Circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func escreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

func main() {
	r := Retangulo{10, 15}
	escreverArea(r) // A área da forma é 150.00

	c := Circulo{10}
	escreverArea(c) // A área da forma é 314.16
}
