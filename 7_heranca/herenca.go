package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Print("herença")
	//var pessoa = pessoa{nome: "renato", sobrenome: "fagalde", idade: 42, altura: 180}
	p1 := estudante{pessoa: pessoa{nome: "renato", sobrenome: "fagalde", idade: 42, altura: 180}, curso: "teste", faculdade: "teste"}
	fmt.Println(p1.pessoa.nome)
}
