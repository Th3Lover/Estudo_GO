package main

import (
	"fmt"
)

func main() {
	//	i := 0

	//	for i < 10 {
	//		i++
	//		fmt.Println("Incrementando i")
	//		time.Sleep(time.Second)
	//	}
	//	fmt.Println(i)

	//	for j := 0; j < 10; j += 2 {
	//		fmt.Println("Incrementando J", j)
	//		time.Sleep(time.Second)
	//	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes { // para não mostrar o indice usar _ no primeiro campo
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario { //Não da para fazer range com struct, apenas com map
		fmt.Println(chave, valor)
	}
}
