package persistence

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql",
		"dingus_user:secret@tcp(127.0.0.1:3306)/dingus")

	if err != nil {
		panic(err)
	}

	return db
}
