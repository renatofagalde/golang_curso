// teste de unidade
package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"
	tipoEsperado := "Avenida"
	tipoNoTeste := TipoDeEndereco(enderecoParaTeste)

	if tipoNoTeste != tipoEsperado {
		t.Error("O tipo recebido não é o esperado")
	}

}
