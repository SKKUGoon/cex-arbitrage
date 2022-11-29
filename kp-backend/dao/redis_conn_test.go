package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"testing"
)

func TestCacheNewConn(t *testing.T) {
	var configFile = "../Redis.yaml"

	redisInfo := map[string]redisLogin{}
	dat, err := os.ReadFile(configFile)
	if err != nil {
		log.Panicln("Conn config file error:", err)
	}
	err = yaml.Unmarshal(dat, &redisInfo)
	if err != nil {
		log.Panicln("Conn config file parse error", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisInfo["conn"].Host + ":" + redisInfo["conn"].Port,
		Password: redisInfo["login"].Password,
		DB:       0,
	})
	// Ping redis client - connection check
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		t.Fatal("database not connected", err)
	}
}
