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
	ws2.TradePubSub = ws2.NewTradeReceiver("./Redis.yaml")
	wsBase := api.New()
	go func() {
		ws2.TradePubSub.Run()
	}()
	wsBase.Conn.GET("/ws", func(context *gin.Context) {
		ws2.WebSocketHandler(context.Writer, context.Request)
	})
	ws := wsBase.Serve("./Config.yaml", *envPtr)
	log.Fatal(ws.ListenAndServe())
}
