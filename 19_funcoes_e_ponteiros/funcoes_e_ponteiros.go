package main

import "fmt"

func inverterSinalComPonteiro(numero *int) int {

	//fmt.Println("numero: ", &numero, " memória: ", numero)

	return *numero * -1
}

func trocarNome(nome *string) {
	fmt.Println("variavel nome: ", *nome, " com endereco de mémoria: ", &nome)
	//valor := *nome
	*nome = *nome + "_teste"
	fmt.Println("variavel nome após troca via endereco de memória: ", *nome, " com endereco de mémoria: ", &nome)
	fmt.Println(*nome)
}

func main() {
	//numero := 20
	//numeroInvert_ido := inverterSinal(numero) //passando uma copia do valor
	//fmt.Println(numeroInvertido)

	numeroNovo := 40
	numeroInvertido01 := inverterSinalComPonteiro(&numeroNovo) //& informando o endereco de mémoria
	fmt.Println("numeroNovo:", numeroNovo, " numeroIvertido: ", numeroInvertido01)

	nome := "renato"
	trocarNome(&nome)
	fmt.Println("nome apos a chamada do ponteiro:", nome)
}
