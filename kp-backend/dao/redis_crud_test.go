package dao

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

var cl *redis.Client = CacheNewConn("../Redis.yaml")

var testData = RdbKeyFieldValue[float64]{
	Key:   "key:uuid",
	Field: "joker",
	Value: 69.2,
}

var testKey = RdbKeyField{
	Key:   "key:2",
	Field: "joker",
}

func TestRdbOpCreate(t *testing.T) {
	RdbOpCreate(cl, ctx, testData)
}

func TestRdbOpRead(t *testing.T) {
	s, err := RdbOpRead(cl, ctx, testKey)
	fmt.Println(s)
	if err != nil {
		t.Fail()
	}
}
