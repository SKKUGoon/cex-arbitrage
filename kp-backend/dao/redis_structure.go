package dao

import "github.com/go-redis/redis/v8"

type redisNewLogin struct {
	Login   redisPI
	Conn    redisConn
	Channel redisChannel
}

type redisPI struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type redisConn struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type redisChannel struct {
	TradeChannel  string `yaml:"trade"`
	SignalChannel string `yaml:"signal"`
}

type RdbKeyFieldValue[T any] struct {
	Key   string
	Field string
	Value T
}

type RdbKeyField struct {
	Key   string
	Field string
}

type Signal[T any] struct {
	Type   string `json:"type" example:"iexa"`
	Status bool   `json:"status" example:"true"`
	Data   T      `json:"data,omitempty"`
}

type Position struct {
	Type     string  `json:"t" example:"enter"`
	Xlong    string  `json:"exl" example:"upbit"`
	Xshort   string  `json:"exs" example:"binance"`
	Asset    string  `json:"a" example:"DOGE"`
	PrcLong  float64 `json:"pl" example:"103"`
	PrcShort float64 `json:"ps" example:"0.07732"`
}

type StatusMessage struct {
	Message string `json:"message" example:"some message"`
}

type CurrentPremium struct {
	ExchangePair struct {
		Long  string `json:"long" example:"upbit"`
		Short string `json:"short" example:"binance"`
	} `json:"exchange_pair"`
	AssetPremium struct {
		Asset           string  `json:"asset" example:"BTC"`
		Premium         float64 `json:"premium" example:"3.50"`
		LongBestAskPrc  float64 `json:"long_ex_bap" example:"17500"`
		ShortBestBidPrc float64 `json:"short_ex_bbp" example:"17500"`
	} `json:"asset_premium"`
}

type SignalMessageQueue struct {
	client        *redis.Client
	SignalMessage chan []byte
	TradeMessage  chan []byte
}
