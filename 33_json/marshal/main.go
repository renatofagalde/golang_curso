package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"handler"
	"log"
)

func main() {
	//json.Marshal()  map -> json
	//json.Unmarshal() json -> map

	java := cachorro.Cachorro{"java", "shitzsu", 5}
	fmt.Println(java)

	javaJSON, erro := json.Marshal(java)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(javaJSON) //slice de bytes, slice de uint8

	//método newBuffer recebe um array de bytes
	fmt.Println(bytes.NewBuffer(javaJSON))

	bagunca := map[string]interface{}{
		"nome":  "bagunca",
		"raca":  "Corge",
		"idade": 2,
	}

	baguncaJSON, erro := json.Marshal(bagunca)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(baguncaJSON))
}
