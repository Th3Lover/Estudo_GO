package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Struct")

	var u usuario
	u.nome = "Davi" //Método padrão
	u.idade = 21
	fmt.Println(u)

	endereco1 := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Pedro", 13, endereco1} // Jeito rapido
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Renato"} //Passando só um valor
	fmt.Println(usuario3)
}
