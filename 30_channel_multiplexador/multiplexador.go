package main

import (
	"fmt"
	"math/rand"
	"time"
)

// a ideia é juntar um ou mais canais e colocar num só
func main() {
	canal := multiplexar(escrever("ola mundo"), escrever("ola mundo, programando em GO"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalEntrada01, canalEntrada02 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canalEntrada01:
				canalSaida <- mensagem

			case mensagem := <-canalEntrada02:
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
