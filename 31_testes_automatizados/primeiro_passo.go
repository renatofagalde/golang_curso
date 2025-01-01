package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

/*
inicializacao do módulo
go mod init introducao-testes
*/
func main() {
	achoTipoEnderenco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println("Achou o tipo de endereco? ", achoTipoEnderenco)
}
