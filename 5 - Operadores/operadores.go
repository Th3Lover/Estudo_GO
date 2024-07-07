package main

import "fmt"

func main() {
	//Operadores aritiméicoos
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 5
	divisao := 10 / 4 //variaceis de tipos diferentes não podem ser usadas junats, int16 com int32
	restodadivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restodadivisao)

	//Operadores de atribuição
	fmt.Println("----------Operadores de atribuição--------")
	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)

	//Operadores relacionais
	fmt.Println("----------Operadores relacionais--------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//Operadores lógicos
	fmt.Println("---------Operadores lógicos---------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Operadores uinários
	fmt.Println("---------Operadores uinários---------")
	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)
	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

}
