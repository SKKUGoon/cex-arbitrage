package main

import (
	"flag"
	"kimchi/api"
	ws2 "kimchi/ws"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Parse Flag
	envPtr := flag.String("env", "dev", "deploy environment")
	flag.Parse()

	// Updater
	var updaterEnv string
	switch *envPtr {
	case "dev":
		updaterEnv = "./Redis.yaml"
	case "deploy":
		updaterEnv = "./Redis_deploy.yaml"
	}
	ws2.TradePubSub = ws2.NewTradeReceiver(updaterEnv)
	wsBase := api.New(updaterEnv)
	go func() {
		ws2.TradePubSub.Run()
	}()

	wsBase.Conn.GET("/ws", func(context *gin.Context) {
		ws2.WebSocketHandler(context.Writer, context.Request)
	})
	ws := wsBase.Serve("./Config.yaml", *envPtr)
	log.Fatal(ws.ListenAndServe())
}
