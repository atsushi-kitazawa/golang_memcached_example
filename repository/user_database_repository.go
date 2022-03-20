package repository

import (
	"database/sql"
	"fmt"

	"github.com/atsushi-kitazawa/golang_memcached_example/database"
	"github.com/atsushi-kitazawa/golang_memcached_example/model"
)

type IUserRepository interface {
	Save(*model.User)
	FindById(string) *model.User
}

type UserRepository struct {
	dbhandler database.IDbHandler
}

func NewUserRepository(hander database.IDbHandler) *UserRepository {
	return &UserRepository{dbhandler: hander}
}

func (u *UserRepository) Save(user model.User) {
	db := u.dbhandler.GetConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction. err=" + err.Error())
		return
	}

	id := user.Id
	name := user.Name
	birthday := user.Birthday
	err = db.QueryRow("INSERT INTO users (id, name, birthday) VALUES ($1, $2, $3)", id, name, birthday).Scan()
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("failed to save user. err=" + err.Error())
		tx.Rollback()
		return
	}
	tx.Commit()

	Set(user.Id, user.UserJsonValue())
}

func (u UserRepository) FindById(id string) *model.User {
	val, err := Get(id)
	if err == nil {
		fmt.Println("cache hit.")
		return model.UserFromJson(val)
	}

	fmt.Println("cache not hit.")
	db := u.dbhandler.GetConnection()
	var user model.User
	err = db.QueryRow("SELECT id, name, birthday FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Birthday)
	if err != nil {
		panic(err)
	}

	return &user
}
