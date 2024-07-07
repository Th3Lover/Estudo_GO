package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Olá mundo")

	var array1 [5]string
	array1[0] = "posição 1"
	fmt.Println(array1)

	array2 := [5]string{"array 2"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{11, 12, 13, 14, 15, 16, 17, 18}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 23) //slice funciona de forma semelhante ao ponteiro
	fmt.Println(slice)        //pode alterar um array livremente

	//aarays internos

	fmt.Println("-----------------")

	slice3 := make([]float32, 10, 11) //define tamanho mino e maximo
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
