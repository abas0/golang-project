package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	//conexao := "root:14051993aA.@tcp(localhost)/minhalivraria"
	db, err := sql.Open("mysql", "root:14051993aA.@/minhalivraria")
	if err != nil {
		panic(err.Error())
	}

	return db
}
