// teste de unidade
package enderecos

import "testing"

type cenarioTest struct {
	enderecoTeste    string
	enderecoEsperado string
}

const naoLocalizado = "Não localizado"

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosTest := []cenarioTest{
		{"Rua Abc", "Rua"},
		{"Avenida apulista", "Avenida"},
		{"Rodovia Abc", "Rodovia"},
		{"Praça das rodas", naoLocalizado},
	}

	for _, cenario := range cenariosTest {
		tipoNoTeste := TipoDeEndereco(cenario.enderecoTeste)
		if tipoNoTeste != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava \"%s\" e recebeu \"%s\"",
				cenario.enderecoEsperado, tipoNoTeste)
		}
	}
}

func TestTipoDeEnderecoErrado(t *testing.T) {
	t.Parallel()

	cenariosTest := []cenarioTest{
		{"Ruas Abc", "Não localizado"},
		{"Avenida apulista", "Avenida"},
		{"Rodovia Abc", "Rodovia"},
		{"Praça das rodas", naoLocalizado},
	}

	for _, cenario := range cenariosTest {
		tipoNoTeste := TipoDeEndereco(cenario.enderecoTeste)
		if tipoNoTeste != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava \"%s\" e recebeu \"%s\"",
				cenario.enderecoEsperado, tipoNoTeste)
		}
	}
}
