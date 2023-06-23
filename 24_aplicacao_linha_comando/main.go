package main

import (
	. "linha-comando/app" //linha de comando é o nome do módulo
	"log"
	"os"
)

func main() {

	aplicacao := Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
