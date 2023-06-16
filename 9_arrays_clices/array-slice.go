package main

import (
	"fmt"
	"reflect"
)

func main() {
	//array
	var arrary1 [5]string
	arrary1[0] = "teste"
	fmt.Println(arrary1)

	//inferencia
	array2 := [5]string{"aryton", "ronaldo", "romario", "zico", "emerson sheik"}
	fmt.Println(array2)

	//outra forma de declarar
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3)

	//slice
	slice := []int{10, 12}
	fmt.Println(slice)

	fmt.Println("arrays: ", reflect.TypeOf(array3))
	fmt.Println("slice: ", reflect.TypeOf(slice))

	//adicionando mais itens
	slice = append(slice, 13)
	fmt.Println(slice)

	//neste exemplo o get(1) é inclusivo, o get(3) é exclusivo, ou seja, ele pega até uma posição anterior
	slice02 := array2[1:3]
	fmt.Println(slice02)

	array2[1] = "ronaldo nazário"
	fmt.Println(array2)

	//arrays internos
	fmt.Println("---------------------")
	slice3 := make([]float32, 3, 4) //slice com o make, sempre dobra a capacidade quando usar um tamanho maior
	fmt.Println(slice3)
	fmt.Println(len(slice3))                       //tamanho
	fmt.Println("slice capacidade: ", cap(slice3)) //capacidade

	//estourando a capacidade do "array", que neste caso é um slice, o mesmo que um array interno
	slice3 = append(slice3, 1)
	slice3 = append(slice3, 1)

	fmt.Println(slice3)
	fmt.Println(len(slice3))                       //tamanho
	fmt.Println("slice capacidade: ", cap(slice3)) //capacidade

}
