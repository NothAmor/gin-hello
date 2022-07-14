package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	db, err := sql.Open("mysql", "gin-hello:wScRz3KRh2xxbRXR@tcp(47.101.38.190:3306)/gin-hello")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	return db
}
