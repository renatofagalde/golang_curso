package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // import implicito
	"log"
)

// Conexao com o banco de dados
// neste retorno apenas um pode existir, pq eles são mutualmente exclusivos.
func Conectar() (*sql.DB, error) {
	stringConexao := "aplicacao:1212@tcp(localhost:3300)/dev?charset=utf8&parseTime=true&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Printf("Erro ao conectar: %s", erro)
		return nil, erro //retorno desta função é exclusivamente
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
