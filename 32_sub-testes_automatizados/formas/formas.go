package formas

import (
	"math"
)

type Formas interface {
	Area() float64
}

// Estrutura de um retangulo
type Retangulo struct {
	altura  float64
	largura float64
}

// estrutura de um circulo
type Circulo struct {
	raio float64
}

// a declaracao não é explicita, ela é implicita
func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

// calcula área de um circulo
func (c Circulo) Area() float64 {
	//return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2)
}
