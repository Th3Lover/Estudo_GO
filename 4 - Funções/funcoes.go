package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func caculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func mult(n1, n2 int8) int8 {
	multi := n1 * n2
	return multi
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	multi := mult(5, 2)
	fmt.Println(multi)

	resultado := f("Texo da função 1")
	fmt.Println(resultado)

	resultadosoma, resultadosubtracao := caculos(10, 15) // o _ serve para chamar uma função e não usar todos os parametros
	fmt.Println(resultadosoma, resultadosubtracao)
}
