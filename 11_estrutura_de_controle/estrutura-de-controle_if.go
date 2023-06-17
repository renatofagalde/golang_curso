package main

import "fmt"

func main() {
	numero := 15
	if numero > 15 {
		fmt.Println("numero maior que 15")
	} else {
		fmt.Println("numero menor que 15")
	}

	fmt.Println("------------------")
	//if init
	if outroNumero := (numero + 10); outroNumero > 15 {
		fmt.Println("numero maior que 15 😊😊")
	} else {
		fmt.Println("numero menor que 15")
	}
}
