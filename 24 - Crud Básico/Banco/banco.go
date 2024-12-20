package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Usando o pacote de fimra implícita
)

// Conectar abre a conexão com o Banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
