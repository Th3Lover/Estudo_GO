package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados do Usuário %s no Banco de Dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 27}
	usuario2.salvar()

	idade := usuario2.maiorDeIdade()
	fmt.Println(idade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
