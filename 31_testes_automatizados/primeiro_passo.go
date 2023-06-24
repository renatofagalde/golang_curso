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
	achoTipoEnderenco := enderecos.TipoEnderenco("Rua tiradentes")
	fmt.Println("Achou o tipo de endereco? %t", achoTipoEnderenco)
}
