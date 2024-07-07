package main

import "fmt"

func main() {

	fmt.Println("Olá mundo")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Ponteiro é uma referencia de memória

	var va1 int
	var pon *int

	va1 = 100
	pon = &va1

	fmt.Println(va1, pon) // Se colocar um * na frente da variavel do poteiro
	// Ele mostra o valor normal, é uma dedsreferenciação

	var va2 int
	var pont *int //ponteiro com valor nulo

	fmt.Println(va2, pont)
}
