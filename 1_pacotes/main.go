package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	auxiliar.EscreverTesteMensagem()
	fmt.Println("escrevendo do arquivo main do modulo")

	erro := checkmail.ValidateFormat(`renato@likwi.com.br`)
	fmt.Println(erro)
	erro01 := checkmail.ValidateFormat(`renato@likwi_com.br`)
	fmt.Println(erro01)
}
