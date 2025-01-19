package redis

import (
	"testing"
)

func TestInitRedis(t *testing.T) {
    InitRedis()

    AddToCache("string", "string", 0) // write tests
}
