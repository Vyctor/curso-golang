package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pow(c.raio, 2) * math.Pi
}

func escreverArea(f forma) {
	fmt.Println("A área da forma é: ", f.area())
}

func main() {
	r := retangulo{10, 15}
	c := circulo{20}

	escreverArea(r)
	escreverArea(c)
}
