package model

import "encoding/json"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func NewUser(id, name, birthday string) *User {
	return &User{
		Id:       id,
		Name:     name,
		Birthday: birthday,
	}
}
func (u *User) UserJsonValue() string {
	e, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return string(e)
}

func UserFromJson(s string) *User {
	var u User
	err := json.Unmarshal([]byte(s), &u)
	if err != nil {
		panic(err)
	}
	return &u
}
