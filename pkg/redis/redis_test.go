package redis

import (
	"testing"
)

func TestInitRedis(t *testing.T) {
    InitRedis()
}

func TestAddToCache(t *testing.T) {
    InitRedis()
    err := AddToCache("key", "value", 60 * 1000 *1000 *1000)

    if err != nil {
        panic(err.Error())
    }

    value, err := GetCache("key")

    if err != nil || value != "value" {
        panic("Unable to get cached value!")
    }

    ClearCache()
}
