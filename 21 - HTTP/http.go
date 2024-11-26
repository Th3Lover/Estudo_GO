package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários!"))
}

func raiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página Raiz!"))
}

func main() {
	//HTTP é um protocolo de comunicação

	//Cliente - Servidor

	//Request - Response

	//Rotas
	//URI - Indentificador de Recursos
	//Método - GET, POST, PUT, DELETE

	http.HandleFunc("/", raiz)

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
