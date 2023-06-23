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

	//map com map
	usuario02 := map[string]map[string]string{
		"pessoa01": {"nome01": "zé ruela", "sobrenome02": "turma da monica"},
		"curso":    {"nome": "culinária", "escola": "diadema cursos online"},
	}

	fmt.Println(usuario02)

	//deletando chave
	delete(usuario02, "curso")
	fmt.Println(usuario02)

	usuario02["signo"] = map[string]string{
		"nome": "virgem",
	}

	fmt.Println(usuario02)
}
