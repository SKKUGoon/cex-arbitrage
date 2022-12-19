package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
)

func TestCacheNewConn(t *testing.T) {
	var configFile = "../config_deploy.yaml"

	redisInfo := map[string]redisNewLogin{}
	dat, err := os.ReadFile(configFile)
	fmt.Println(dat)
	if err != nil {
		log.Panicln("Conn config file error:", err)
	}

	err = yaml.Unmarshal(dat, &redisInfo)
	fmt.Println(redisInfo)
	if err != nil {
		log.Panicln("Conn config file parse error", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisInfo["conn"].Conn.Host + ":" + redisInfo["conn"].Conn.Port,
		Password: redisInfo["login"].Login.Password,
		DB:       0,
	})
	// Ping redis client - connection check
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		t.Fatal("database not connected", err)
	}
}
