package main

import "fmt"

func main() {
	variavel01, variavel02 := 10, 10
	fmt.Println(variavel01, variavel02)

	variavel01++
	fmt.Println(variavel01, variavel02)

	//ponteiro é uma referencia na memória
	variavel03 := 10
	var ponteiro *int

	fmt.Println(variavel03, ponteiro)
	ponteiro = &variavel03 //declaracao
	variavel03++

	//apos a troca de valores
	fmt.Println(variavel03, ponteiro)

	//resgatar o valor através do ponteiro:
	//desreferenciação
	variavel04 := *ponteiro
	variavel04++
	fmt.Println(variavel03, variavel04)
}
