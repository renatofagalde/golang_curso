package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculo(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(texto string) string {
		return texto
	}
	fmt.Println(f("texto dentro da funcao"))

	soma01, subtracao01 := calculo(10, 20)
	fmt.Println("Valores retornados em multiplas funcoes soma: %d, subtracao01: %d", soma01, subtracao01)

	//o underline significa que é para ignorar um dos retornos,
	//lembrando que isto é obrigatório pq go não compila com uma variável
	//que não está sendo usada
	_, subtracao02 := calculo(10, 20)
	fmt.Println("Valor da subtracao ignorando a soma ", subtracao02)
}
