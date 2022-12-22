package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func handleMsg(conn *websocket.Conn, wd chan struct{}) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("type", msgType)
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", msg)
	}
}

func handleSubscription(conn *websocket.Conn) {
	s, _ := json.Marshal(TestBinanceSubscribe)
	fmt.Println(string(s))
	err := conn.WriteJSON(TestBinanceSubscribe)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("subscription sent")
}

func Do(conn *websocket.Conn, wd chan struct{}, osSig chan os.Signal) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-wd:
			return
		case <-osSig:
			log.Println("interrupted by ctrl+c")
			err := conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(
					websocket.CloseNormalClosure,
					"",
				),
			)
			if err != nil {
				log.Println("err write close:", err)
				return
			}
			return
		case <-ticker.C:

		}
	}
}
