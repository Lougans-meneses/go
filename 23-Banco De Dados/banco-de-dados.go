package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// string de conexao para o banco MYSQL
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	// fechando para não ficar aberto
	defer db.Close()

	// Ping() vai fazer um teste de conexao com o banco 
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}

	// fechando para não ficar aberto
	defer linhas.Close()

	fmt.Println(linhas)
}