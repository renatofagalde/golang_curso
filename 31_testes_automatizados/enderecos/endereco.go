package enderecos

import "strings"

// TipoEnderenco localizar o tipo do endereco
func TipoEnderenco(endereco string) bool {
	tipos := []string{"rua", "avenida", "estrada", "rodivia"}
	primeiraPalavra := strings.Split(strings.ToLower(endereco), " ")[0]

	for _, tipo := range tipos {
		if tipo == primeiraPalavra {
			return true
		}
	}
	return false
}
