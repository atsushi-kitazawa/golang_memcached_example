package model

import "encoding/json"

type user struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func UserJsonValue(id, name, birthday string) string {
	user := &user{
		Id:       id,
		Name:     name,
		Birthday: birthday,
	}
	e, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return string(e)
}
