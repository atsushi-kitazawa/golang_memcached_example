package repository

import (
	"database/sql"

	"github.com/atsushi-kitazawa/golang_memcached_example/model"

	_ "github.com/lib/pq"
)

func connection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=35432 user=postgres password=password dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func Save(user model.User) {
	db := connection()
	defer db.Close()

	id := user.Id
	name := user.Name
	birthday := user.Birthday
	db.QueryRow("INSERT INTO users (id, name, birthday) VALUES ($1, $2, $3)", id, name, birthday)
}

func FindById(id string) *model.User {
	db := connection()

	var user model.User
	err := db.QueryRow("SELECT id, name, birthday FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Birthday)
	if err != nil {
		panic(err)
	}

	return &user
}
