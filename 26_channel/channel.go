package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) //chan palavra reservada para criar canal
	go escrever("olá mundo", canal)

	/*	for {
			mensagem, canalAberto := <-canal
			if !canalAberto {
				break
			}
			fmt.Println(mensagem)
		}
	*/
	//usando de forma simplificada
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal) //fechando o canal
}
