package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

var cl *redis.Client = CacheNewConn("../Redis.yaml")
var ctx = context.Background()

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
	RdbOpCreate[float64](cl, ctx, testData)
}

func TestRdbOpRead(t *testing.T) {
	s, err := RdbOpRead(cl, ctx, testKey)
	fmt.Println(s)
	if err != nil {
		t.Fail()
	}
}
