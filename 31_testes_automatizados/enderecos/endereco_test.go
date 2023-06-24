// teste de unidade
package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenide Paulista"
	tipoEsperado := "Avenida"
	tipoNoTeste := TipoDeEndereco(enderecoParaTeste)

	if tipoNoTeste != tipoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava \"%s\" e recebeu \"%s\"",
			tipoEsperado, tipoNoTeste)
	}

}
