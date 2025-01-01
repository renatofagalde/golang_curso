package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(true)

	//exemplo classico é fmt.Println

	//exemplo de qualquer coisa:
	mapa := map[interface{}]interface{}{
		1:                "exemplo",
		float64(1234345): 1111,
		true:             12345,
	}

	fmt.Println(mapa)
}
