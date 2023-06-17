package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalizador01 := soma(1, 2, 3, 4, 5, 1, 2, 4, 5, 6, 1)
	fmt.Println(totalizador01)
}
