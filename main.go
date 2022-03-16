package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	doMain()
}

func doMain() {
	mc := memcache.New("localhost:11211")
	mc.Set((&memcache.Item{Key: "foo", Value: []byte("value")}))

	it, err := mc.Get("foo")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(it.Value))
}
