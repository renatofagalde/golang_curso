package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("salvando usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// por conta do * isto é uma referencia
// por isto ele troca
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario := usuario{"john doe", 20}
	usuario.salvar()

	fmt.Printf("usuario %s é maior de idade? %v\n", usuario.nome, usuario.maiorDeIdade())

	usuario.fazerAniversario()

	fmt.Println(usuario)
}
