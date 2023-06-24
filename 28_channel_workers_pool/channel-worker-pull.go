package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	/*	go worker(tarefas, resultados)
		go worker(tarefas, resultados)
		go worker(tarefas, resultados)
		go worker(tarefas, resultados)
		go worker(tarefas, resultados)
		go worker(tarefas, resultados)
		go worker(tarefas, resultados)
	*/
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	elapsed := time.Since(start)
	log.Printf("\n45 fibonacci took %s", elapsed)
}
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}
func fibonacci(posicao int) int {
	// usar o memoization

	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
