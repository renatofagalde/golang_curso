package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import implicito
	"log"
)

func main() {
	stringConexao := "aplicacao:1212@tcp(localhost:3300)/dev?charset=utf8&parseTime=true&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	//validar conexao com o Ping
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("conexão aberta")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
