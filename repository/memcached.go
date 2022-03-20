package repository

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var mc *memcache.Client

func init() {
	mc = memcache.New("localhost:11211", "localhost:11212")
}

func Set(key, value string) {
	mc.Set(&memcache.Item{Key: key, Value: []byte(value)})
}

func Get(key string) (string, error) {
	item, err := mc.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}
