package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //Usando o pacote de fimra implícita
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local" // "usuário:senha@/nome do banco"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println("Erro no Open")
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		fmt.Println("Erro no Ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
