package main

import "fmt"

func main() {
	var numero int8 = 100
	fmt.Println(numero)

	//alias INT32 = RUNE
	var numero2 rune = 12345
	fmt.Println(numero2)

	//byte = int8
	var numero3 byte = 123
	fmt.Println(numero3)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 = 1222222.45
	fmt.Println(numeroReal2)

	var str string = "texto"
	fmt.Println(str)
}
