package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"]) //para acessar um compo expecifico

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joão",
			"ultimo":   "pedro",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "SJC",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}

	fmt.Println(usuario2)
}
