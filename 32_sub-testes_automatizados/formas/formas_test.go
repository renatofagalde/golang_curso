package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := retangulo.Area()

		if areaEsperada != areaRecebida {
			// pode ser usado o Fatal, neste caso ele interrompe.
			t.Fatalf("Área recebida %f é diferente de esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {

		circulo := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			// pode ser usado o Fatal, neste caso ele interrompe.
			t.Fatalf("Área recebida %f é diferente de esperada %f", areaRecebida, areaEsperada)
		}
	})
}
