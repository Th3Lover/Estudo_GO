package main

import (
	"fmt"
	"math"
)

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

type forma interface {
	area() float64
}

func EscreverArea(f forma) {
	fmt.Println("A área da forma é:", f.area())
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	EscreverArea(r)

	c := circulo{10}
	EscreverArea(c)
}
