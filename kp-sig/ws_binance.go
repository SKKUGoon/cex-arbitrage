package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func handleBinanceMsg(conn *websocket.Conn, wd chan struct{}) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("type", msgType)
			log.Println("msg", err.Error())
		}
		log.Printf("recv: %s", msg)
	}
}

func sendBinanceSub(subscr BinanceSubscribe, conn *websocket.Conn) {
	err := conn.WriteJSON(subscr)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Binance Subscription Sent")
}
