package main

import "fmt"

func main() {
	var variavel1 string = "teste1"
	fmt.Println(variavel1)

	variavel2 := "teste2" //declaracao por inferencia
	fmt.Println(variavel2)

	variavel3, variavel4 := "valor por inferência3", " por inferencia 4"
	fmt.Println(variavel3, variavel4)
}
