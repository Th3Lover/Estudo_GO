package main

import "fmt"

func main() {

	fmt.Println("Estruturas")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("Numero maior que zero")
	} else {
		fmt.Println("Numero menor que zero")
	}
}
