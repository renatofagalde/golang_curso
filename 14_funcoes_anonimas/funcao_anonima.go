package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println("ola mundo ", texto)
	}("aqui da terra") //funcao anonima

}
