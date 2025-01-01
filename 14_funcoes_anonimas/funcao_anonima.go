package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println("ola mundo ", texto)
	}("aqui da terra") //funcao anonima

	retorno := func(texto string) string {
		return fmt.Sprintf("Recbido o texto: %s", texto)
	}("sou anonimo")

	fmt.Println(retorno)

}
