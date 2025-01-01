package main

import "fmt"

func closure() func() {
	texto := "dentro da funcao closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

// manter o scope das declaracoes de variáveis
func main() {
	texto := "dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
