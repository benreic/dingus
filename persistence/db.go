package persistence

import (
	"database/sql"
	"github.com/benreic/dingus/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewDb() *sql.DB {

	credentials := config.DbUsername() + ":" + config.DbPassword()

	db, err := sql.Open("mysql",
		credentials+"@tcp(127.0.0.1:3306)/dingus")

	if err != nil {
		panic(err)
	}

	return db
}
