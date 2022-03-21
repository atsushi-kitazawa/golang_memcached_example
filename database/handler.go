package database

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDbHandler interface {
	GetDatabase() *gorm.DB
}

type dbhandler struct {
}

func NewDbHandler() IDbHandler {
	return &dbhandler{}
}

func (h *dbhandler) GetDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=testdb port=35432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
