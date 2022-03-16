package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

var mc *memcache.Client

func init() {
	mc = memcache.New("localhost:11211")
}

func main() {
	doMain()
}

func doMain() {
	set("key1", "val1")
	fmt.Println(string(get("key1")))
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
