package repository

import (
	"database/sql"
	"fmt"

	"github.com/atsushi-kitazawa/golang_memcached_example/model"

	_ "github.com/lib/pq"
)

type IUserRepository interface {
	Save(*model.User)
	FindById(string) *model.User
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func getConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=35432 user=postgres password=password dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func (u *UserRepository) Save(user model.User) {
	db := getConnection()
	defer db.Close()

	id := user.Id
	name := user.Name
	birthday := user.Birthday
	db.QueryRow("INSERT INTO users (id, name, birthday) VALUES ($1, $2, $3)", id, name, birthday)

	Set(user.Id, user.UserJsonValue())
}

func (u UserRepository) FindById(id string) *model.User {
	val, err := Get(id)
	if err == nil {
		fmt.Println("cache hit.")
		return model.UserFromJson(val)
	}

	fmt.Println("cache not hit.")
	db := getConnection()
	var user model.User
	err = db.QueryRow("SELECT id, name, birthday FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Birthday)
	if err != nil {
		panic(err)
	}

	return &user
}
