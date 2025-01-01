package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("incrementando i: ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("final i: ", i)

	fmt.Println("---------------")
	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j:", j)
	}
	//j fora do escopo
	//fmt.Println("final j: ", j)

	nomes := [...]string{"ayrton", "ronaldo", "romário"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("---------------")
	for _, nome := range nomes {
		fmt.Println(nome)
	}
	fmt.Println("---------------")

	for indice, letra := range "palavra" {
		fmt.Println(indice, letra, string(letra))
	}
	fmt.Println("---------------")

	//range não funciona em struct
	usuario := map[string]string{
		"nome":      "ayrton",
		"sobrenome": "silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	fmt.Println("---------------")

	//loop infinito
	//for {
	//
	//}
}
