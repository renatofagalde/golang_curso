package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	//especificar a quantidade correta
	waitGroup.Add(2)

	//transformando uma funcao anonima em go routine
	go func() {
		escrever01("ola mundo 1")
		waitGroup.Done()
	}()

	//transformando uma funcao anonima em go routine
	go func() {
		escrever01("programando em go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever01(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
