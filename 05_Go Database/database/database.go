package database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:Kozato321@tcp(localhost:3306)/golearn")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
