package main

import "fmt"

func main() {
	canal := escrever("olá mundo")
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// retornar um canal de string
func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido %s", texto)
		}
	}()

	return canal
}
