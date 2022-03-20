package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type IDbHandler interface {
	GetConnection() *sql.DB
}

type DbHandler struct {
}

func NewDbHandler() *DbHandler {
	return &DbHandler{}
}

func (handler DbHandler) GetConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=35432 user=postgres password=password dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
