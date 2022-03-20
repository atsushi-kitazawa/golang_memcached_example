package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atsushi-kitazawa/golang_memcached_example/model"
	"github.com/atsushi-kitazawa/golang_memcached_example/repository"
)

func main() {
	doMain()
}

func doMain() {
	setCmd := flag.NewFlagSet("set", flag.ExitOnError)
	setKey := setCmd.String("key", "", "key")
	setName := setCmd.String("name", "", "user name")
	setBirthday := setCmd.String("birthday", "", "user birthday")

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getKey := getCmd.String("key", "", "key")

	switch os.Args[1] {
	case "set":
		setCmd.Parse(os.Args[2:])
		user := model.NewUser(*setKey, *setName, *setBirthday)
		repo := repository.NewUserRepository()
		repo.Save(*user)
		fmt.Printf("key %s set value %s\n", *setKey, user)
	case "get":
		getCmd.Parse(os.Args[2:])
		user := repository.NewUserRepository().FindById(*getKey)
		fmt.Printf("key %s values is %s\n", *getKey, user)
	}
}
