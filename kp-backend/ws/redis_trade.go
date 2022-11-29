package ws

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"kimchi/common"
	"kimchi/dao"
)

type TradeReceiver struct {
	client *redis.Client

	Messages       chan []byte
	NewConnections chan *websocket.Conn
	RmConnections  chan *websocket.Conn
}

var TradePubSub TradeReceiver

func NewTradeReceiver(configFile string) TradeReceiver {
	c := dao.CacheNewConn(configFile)
	common.PrintGreenOk("Create Redis connection for websocket")
	return TradeReceiver{
		client:         c,
		Messages:       make(chan []byte, 1000),
		NewConnections: make(chan *websocket.Conn),
		RmConnections:  make(chan *websocket.Conn),
	}
}

func (tr *TradeReceiver) broadcast(msg []byte) {
	tr.Messages <- msg
}

func (tr *TradeReceiver) register(conn *websocket.Conn) {
	common.PrintYellowOperation("wss:: registering connection")
	tr.NewConnections <- conn
}

func (tr *TradeReceiver) deregister(conn *websocket.Conn) {
	common.PrintYellowOperation("wss:: deregistering connection")
	tr.RmConnections <- conn
}

func removeConn(conns []*websocket.Conn, remove *websocket.Conn) []*websocket.Conn {
	var i int
	var found bool
	for i = 0; i < len(conns); i++ {
		if conns[i] == remove {
			found = true
			break
		}
	}
	if !found {
		msg := fmt.Sprintf("conns: %#v\nconn: %#v\n", conns, remove)
		common.PrintPurpleWarning(msg)
		panic("Conn not found")
	}
	copy(conns[i:], conns[i+1:]) // shift down
	conns[len(conns)-1] = nil    // nil last element
	return conns[:len(conns)-1]  // truncate slice
}

func (tr *TradeReceiver) Run() error {
	ctx := context.Background()
	go tr.connHandler()
	subscriber := tr.client.Subscribe(ctx, "trade_channel")
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		tr.broadcast([]byte(msg.Payload))
	}
}

func (tr *TradeReceiver) connHandler() {
	conns := make([]*websocket.Conn, 0)
	for {
		select {
		case msg := <-tr.Messages:
			for _, conn := range conns {
				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					common.PrintPurpleWarning("wss:: error writing data. closing and removing connections")
					conns = removeConn(conns, conn)
				}
			}
		case conn := <-tr.NewConnections:
			conns = append(conns, conn)
			common.PrintCyanStatus(fmt.Sprintf("wss:: current connection #: %v", len(conns)))
		case conn := <-tr.RmConnections:
			conns = removeConn(conns, conn)
			common.PrintCyanStatus(fmt.Sprintf("wss:: current connection #: %v", len(conns)))
		}
	}
}
