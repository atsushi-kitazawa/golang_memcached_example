package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atsushi-kitazawa/golang_memcached_example/model"
	"github.com/bradfitz/gomemcache/memcache"
)

var mc *memcache.Client

func init() {
	mc = memcache.New("localhost:11211", "localhost:11212")
}

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
		user := model.UserJsonValue(*setKey, *setName, *setBirthday)
		set(*setKey, user)
		fmt.Printf("key %s set value %s\n", *setKey, user)
	case "get":
		getCmd.Parse(os.Args[2:])
		val := string(get(*getKey))
		fmt.Printf("key %s values is %s\n", *getKey, val)
	}
}

func set(key, value string) {
	mc.Set(&memcache.Item{Key: key, Value: []byte(value)})
}

func get(key string) []byte {
	item, err := mc.Get(key)
	if err != nil {
		panic(err)
	}
	return item.Value
}
