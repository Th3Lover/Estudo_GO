package main

import "fmt"

func fibonaci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonaci(posicao-2) + fibonaci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(12)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonaci(i))
	}
}
