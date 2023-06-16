package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	var u01 usuario
	u01.nome = "teste"
	u01.idade = 10

	//inferencia
	u02 := usuario{"usuario01", 10}
	fmt.Println(u02)

	u03 := usuario{idade: 21}
	fmt.Println(u03)
}
