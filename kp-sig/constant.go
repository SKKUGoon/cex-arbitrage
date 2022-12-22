package main

import (
	"fmt"
	"strings"
)

const BINANCE = "wss://stream.binance.com:9443/ws"
const UPBIT = "wss://api.upbit.com/websocket/v1"

type BinanceSubscribe struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

type UpbitSubscribe struct {
	// First part
	Ticket string `json:"ticket,omitempty"`
	// Second part
	Type           string   `json:"type,omitempty"`
	Codes          []string `json:"codes,omitempty"`
	IsOnlyRealTime bool     `json:"isOnlyRealtime,omitempty"`
}

type BinanceBookTicker struct {
	LogTime int    `json:"u"`
	Ticker  string `json:"s"`
	Bid     string `json:"b"`
	BidQty  string `json:"B"`
	Ask     string `json:"a"`
	AskQty  string `json:"A"`
}

var TestBinanceSubscribe = BinanceSubscribe{
	Method: "SUBSCRIBE",
	Params: []string{"linkbusd@bookTicker"},
	ID:     42,
}

var TestUpbitSubscribe1 = UpbitSubscribe{
	Ticket: "test",
}
var TestUpbitSubscribe2 = UpbitSubscribe{
	Type:           "ticker",
	Codes:          []string{"KRW-BTC"},
	IsOnlyRealTime: true,
}

var TestUpbitSubscribe = []UpbitSubscribe{
	TestUpbitSubscribe1,
	TestUpbitSubscribe2,
}

func genBinanceSub(asset []string, keyCurrency string, subType bool) BinanceSubscribe {
	subAsset := []string{}
	for _, a := range asset {
		m := fmt.Sprintf(
			"%s%s@bookTicker",
			strings.ToLower(a),
			strings.ToLower(keyCurrency),
		)
		subAsset = append(subAsset, m)
	}
	switch subType {
	case true:
		return BinanceSubscribe{
			Method: "SUBSCRIBE",
			Params: subAsset,
			ID:     42,
		}
	default:
		return BinanceSubscribe{
			Method: "UNSUBSCRIBE",
			Params: subAsset,
			ID:     42,
		}
	}
}

// func genUpbitSub() {
// 	subAsset := []string{}
// 	for _, a := range asset {
// 		m := fmt.Sprintf(
// 			%s%
// 		)
// 	}
// }
