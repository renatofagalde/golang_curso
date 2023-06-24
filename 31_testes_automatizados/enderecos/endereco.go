package enderecos

import "strings"

// TipoDeEndereco localizar o tipo do endereco
func TipoDeEndereco(endereco string) string {
	tipos := []string{"rua", "avenida", "estrada", "rodivia"}
	primeiraPalavra := strings.Split(strings.ToLower(endereco), " ")[0]
	achou := false
	for _, tipo := range tipos {
		if tipo == primeiraPalavra {
			achou = true
			break
		}
	}

	if achou {
		return strings.Title(primeiraPalavra)
	}

	return "Não localizado"
}
