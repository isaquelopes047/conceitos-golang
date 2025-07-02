package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "root:A52-FAdCc-BB4EGAEE4GacaEggF-Ahb4@tcp(roundhouse.proxy.rlwy.net:44062)/railway?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão esta aberta!")

	linhas, erro := db.Query("SELECT * FROM AvisosGeral")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)
}
