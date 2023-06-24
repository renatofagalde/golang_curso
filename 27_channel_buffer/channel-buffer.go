package main

import "fmt"

/*
//deadlock porque o canal atingiu o número de buffer
func main() {
	canal := make(chan string)
	canal <- "olá"

	mensagem := <-canal
	fmt.Println(mensagem)
}
*/

func main() {
	canal := make(chan string, 2) //liberando o buffer do canal para processar
	canal <- "olá"

	mensagem := <-canal
	fmt.Println(mensagem)
}
