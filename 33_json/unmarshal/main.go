package main

import (
	"encoding/json"
	"fmt"
	cachorro "handler"
	"log"
)

func main() {
	baguncaJSON := `{"idade":2,"nome":"bagunca","raca":"Corge"}`
	javaJSON := `{"idade":5,"nome":"Java","raca":"Shiztsu"}`
	var bagunca cachorro.Cachorro

	// []byte -> isso é um slice de bytes
	//o método unmarshal recebe um slice de bytes, por isso o códio abaixo
	// &bagunca = endereco de memória, ou seja, o que for alteredo no método será
	// refletido aqui neste contexto
	if erro := json.Unmarshal([]byte(baguncaJSON), &bagunca); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bagunca)

	// usando map
	java := make(map[string]interface{})
	if erro := json.Unmarshal([]byte(javaJSON), &java); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(java)

}
