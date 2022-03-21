package repository

import (
	"fmt"

	"github.com/atsushi-kitazawa/golang_memcached_example/database"
	"github.com/atsushi-kitazawa/golang_memcached_example/model"
	"gorm.io/gorm"
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
	db := u.dbhandler.GetDatabase()
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			fmt.Println("failed to save user." + err.Error())
			return err
		}
		return nil
	})

	Set(user.Id, user.UserJsonValue())
}

func (u UserRepository) FindById(id string) *model.User {
	val, err := Get(id)
	if err == nil {
		fmt.Println("cache hit.")
		return model.UserFromJson(val)
	}

	db := u.dbhandler.GetDatabase()
	var user model.User
	db.First(&user, id)

	return &user
}
