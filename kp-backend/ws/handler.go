package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"kimchi/common"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		common.PrintPurpleWarning("wss:: failed to set websocket upgrade: %v", err.Error())
		return
	}

	TradePubSub.register(wsConn)

	for {
		t, msg, err := wsConn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway) || err == io.EOF {
				common.PrintCyanStatus("wss: websocket closed!")
				break
			}
			common.PrintPurpleWarning("wss:: failed to read websocket message")
			break
		}

		switch t {
		case websocket.TextMessage:
			common.PrintCyanStatus(fmt.Sprintf("wss:: inbound message %s", string(msg)))
		default:
			common.PrintPurpleWarning("wss:: inbound unknown message")
		}
	}
	TradePubSub.deregister(wsConn)
	wsConn.WriteMessage(websocket.CloseMessage, []byte{})
}
