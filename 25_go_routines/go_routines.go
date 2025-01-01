package main

import (
	"fmt"
	"time"
)

func main() {
	//concorrencia != paralelismo
	go escrever("ola mundo 1")
	go escrever("programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
