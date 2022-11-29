package main

import (
	"github.com/gin-gonic/gin"
	"kimchi/api"
	ws2 "kimchi/ws"
)

func main() {
	// Updater
	ws2.TradePubSub = ws2.NewTradeReceiver("./Redis.yaml")
	wsBase := api.New()
	go func() {
		ws2.TradePubSub.Run()
	}()
	wsBase.Conn.GET("/ws", func(context *gin.Context) {
		ws2.WebSocketHandler(context.Writer, context.Request)
	})
	ws := wsBase.Serve("./Config.yaml")
	ws.ListenAndServe()
}
