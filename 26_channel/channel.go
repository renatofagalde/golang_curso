package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) //chan palavra reservada para criar canal
	go escrever("olá mundo", canal)
	fmt.Println(<-canal)
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
}
