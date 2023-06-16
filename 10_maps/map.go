package main

import "fmt"

func main() {

	//map[string]string
	//o primeiro string é o tipo da chave
	//o segundo string é o tipo do dado
	usuario := map[string]string{
		"nome":      "ayrton",
		"sobrenome": "silva",
	}

	fmt.Println(usuario)
}
