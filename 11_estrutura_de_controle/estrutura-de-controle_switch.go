package main

import (
	"fmt"
	"math/rand"
	"time"
)

func diaDaSemana(numero int) string {
	fmt.Println("calculando para o dia: ", numero)

	switch numero {
	case 1:
		//fallthrough passa para dentro do codigo da proxima condição, ele não testa, ele apenas passa
		return "domingo"
	case 2:
		return "segunda"
	default:
		return "férias"
	}
}
func main() {
	fmt.Println(diaDaSemana(rand.Intn(5)))

	fmt.Println(time.Now().Weekday().String())
}
