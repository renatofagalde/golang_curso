package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("salvando usuário %s no banco de dados", u.nome)
}

func main() {
	usuario := usuario{"john doe", 10}
	usuario.salvar()

}
