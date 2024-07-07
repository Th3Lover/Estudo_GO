package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa           //não se expecifica o tipo
	curso     string //aqui esta o mais proximo de herança em go
	faculdade string
}

func main() {
	fmt.Println("Teste")

	p1 := pessoa{"Leonardo", "Fernandes", 19, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "Unip"}
	fmt.Println(e1)

	fmt.Println(e1.nome) //puxando apenas o nome
}
