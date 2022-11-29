package api

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"kimchi/dao"
	"testing"
)

var cl *redis.Client = dao.CacheNewConn("../Redis.yaml")
var ctx = context.Background()

func TestGetComparison(t *testing.T) {
	// Get Band information
	var bandMap map[string]string
	var err error

	bandUD := "lower"
	switch bandUD {
	case "upper":
		searchKeyUpper := dao.RdbKeyField{Key: "band_upper"}
		bandMap, err = dao.RdbOpRead(cl, ctx, searchKeyUpper)
	case "lower":
		searchKeyLower := dao.RdbKeyField{Key: "band_lower"}
		bandMap, err = dao.RdbOpRead(cl, ctx, searchKeyLower)
	}
	if err != nil {
		t.Fail()
	}
	fmt.Println(bandMap)
}
