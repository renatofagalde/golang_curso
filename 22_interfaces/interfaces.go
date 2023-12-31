package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("a área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

// a declaracao não é explicita, ela é implicita
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	//return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	fmt.Println("--------")
	c := circulo{10}
	escreverArea(c)
}
