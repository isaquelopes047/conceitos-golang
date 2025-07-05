package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Drive de conexão com o mysql
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "root:A52-FAdCc-BB4EGAEE4GacaEggF-Ahb4@tcp(roundhouse.proxy.rlwy.net:44062)/railway?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
