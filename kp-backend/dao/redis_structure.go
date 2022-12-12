package dao

import "github.com/go-redis/redis/v8"

type redisLogin struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
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

type StatusMessage struct {
	Message string `json:"message" example:"some message"`
}

type Trade struct {
	Exchange string `json:"ex" example:"binance"`
	Position string `json:"p" example:"short"`
	Asset    string `json:"a" example:"BTC/USDT"`
}

type CurrentPremium struct {
	ExchangePair struct {
		Long  string `json:"long" example:"upbit"`
		Short string `json:"short" example:"binance"`
	} `json:"exchange_pair"`
	AssetPremium struct {
		Asset   string  `json:"asset" example:"BTC"`
		Premium float64 `json:"premium" example:"3.50"`
	} `json:"asset_premium"`
}

type SignalMessageQueue struct {
	client        *redis.Client
	SignalMessage chan []byte
	TradeMessage  chan []byte
}
