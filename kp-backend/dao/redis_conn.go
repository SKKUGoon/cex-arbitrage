package dao

import (
	"context"
	"kimchi/common"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
)

func CacheNewConn(configFile string) *redis.Client {
	// configFile should be "../Redis.yaml"
	redisInfo := map[string]redisLogin{}
	dat, err := os.ReadFile(configFile)
	if err != nil {
		log.Panicln("Redis conn config file error:", err)
	}
	err = yaml.Unmarshal(dat, &redisInfo)
	if err != nil {
		log.Panicln("Redis conn config file parse error:", err)
	}

	common.PrintBlueStatus(redisInfo["conn"].Host + ":" + redisInfo["conn"].Port)
	common.PrintBlueStatus(redisInfo["login"].Password)
	client := redis.NewClient(&redis.Options{
		Addr:     redisInfo["conn"].Host + ":" + redisInfo["conn"].Port,
		Password: redisInfo["login"].Password,
		DB:       0,
	})
	// Ping redis client - connection check
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Panicln("database not connected")
	}
	return client
}
