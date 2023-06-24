// teste de unidade
package enderecos

import "testing"

type cenarioTest struct {
	enderecoTeste    string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosTest := []cenarioTest{
		{"Rua Abc", "Rua"},
		{"Avenida apulista", "Avenida"},
		{"Rodovia Abc", "Rodovia"},
		{"Praça das rodas", "Não localizado"},
	}

	for _, cenario := range cenariosTest {
		tipoNoTeste := TipoDeEndereco(cenario.enderecoTeste)
		if tipoNoTeste != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava \"%s\" e recebeu \"%s\"",
				cenario.enderecoEsperado, tipoNoTeste)
		}

	}

}
