package main

import (
	"errors"
	"fmt"
)

// tipos de int no go int8, int16, int32, int64 muda o peso e tamanho de armazenamento

func main() {
	numero := 1000000000000
	fmt.Println(numero)

	// uint -> unsygned int = int sem sinal

	var numero2 uint = 10000
	fmt.Println(numero2)

	//alias -> apelido de referencia

	var numero3 rune = 123456 //apelido que representa exatamente o int32
	fmt.Println(numero3)

	var numero4 byte = 123 //byte = uint8
	fmt.Println(numero4)

	//___________________ FLOAT ________________________________

	var numeroreal1 float32 = 123.45
	fmt.Println(numeroreal1)

	var numeroreal2 float64 = 123000000.45
	fmt.Println(numeroreal2)

	//______________________STRING__________________________

	var str string = "Texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char) // devolve o numero na tabela assci

	//__________________ VALOR ZERO____________________

	var texto int16
	fmt.Println(texto) // como string entrega nada mas se colcoar int ele devolve zero, sendo o valor inicial

	//________________ BOOLEANO______________________

	var booleano1 bool
	fmt.Println(booleano1)

	//__________________ ERROR ________________________

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro Interno")
	fmt.Println(erro2)
}
