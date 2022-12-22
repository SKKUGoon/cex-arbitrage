package main

import (
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	// messageOut := make(chan string)
	interrupt := make(chan os.Signal, 1)

	c, resp, err := websocket.DefaultDialer.Dial(BINANCE, nil)
	if err != nil {
		log.Printf("handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}
	if resp != nil {
		log.Printf("status: %s", resp.Status)
	}
	defer c.Close()

	done := make(chan struct{})
	defer close(done)
	go handleSubscription(c)
	go handleMsg(c, done)

	Do(c, done, interrupt)
}
