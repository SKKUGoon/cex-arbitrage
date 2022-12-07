package api

import (
	"context"
	"errors"
	"fmt"
	"kimchi/common"
	"kimchi/dao"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Signal[T any] struct {
	Type string `json:"type" example:"iexa"`
	Data T      `json:"data"`
}

type CurrentPremium struct {
	ExchangePair struct {
		Long  string `json:"long" example:"upbit"`
		Short string `json:"short" example:"binance"`
	} `json:"exchange_pair"`
	AssetPremium []struct {
		Asset   string  `json:"asset" example:"BTC"`
		Premium float64 `json:"premium" example:"3.50"`
	} `json:"asset_premium"`
}

type Band struct {
	Asset string  `json:"asset" example:"BTC"`
	Upper float64 `json:"upper" example:"3.42"`
	Lower float64 `json:"lower" example:"0.56"`
}

type Trade struct {
	Exchange string `json:"ex" example:"binance"`
	Position string `json:"p" example:"short"`
	Asset    string `json:"a" example:"BTC/USDT"`
}

type StatusMessage struct {
	Message string `json:"message" example:"some message"`
}

func getComparison(bandUD string, client *redis.Client) (map[string]string, error) {
	// Get Band information
	var bandMap map[string]string
	var err error
	switch bandUD {
	case "upper":
		searchKeyUpper := dao.RdbKeyField{Key: "band_upper"}
		bandMap, err = dao.RdbOpRead(client, context.Background(), searchKeyUpper)
	case "lower":
		searchKeyLower := dao.RdbKeyField{Key: "band_lower"}
		bandMap, err = dao.RdbOpRead(client, context.Background(), searchKeyLower)
	}
	if err != nil {
		return nil, err
	}
	return bandMap, nil
}

func comparePremium(p CurrentPremium, client *redis.Client) ([]Signal[Trade], error) {
	common.PrintBlueStatus("Premium Status Comparison")
	bandInfoLow, err1 := getComparison("lower", client)
	bandInfoUp, err2 := getComparison("upper", client)
	if err1 != nil || err2 != nil {
		return []Signal[Trade]{}, errors.New(err1.Error() + err2.Error())
	}

	var compareResult []Signal[Trade]
	for _, premium := range p.AssetPremium {
		thresLow, _ := strconv.ParseFloat(bandInfoLow[premium.Asset], 64)
		thresUp, _ := strconv.ParseFloat(bandInfoUp[premium.Asset], 64)

		if thresUp-thresLow < 0.015 {
			// If the distance between band's upper threshold and lower threshold
			// are so small, arbitrage might not be possible. Therefore, set a
			// minumum threshold level and if the distance is not large enough
			// return false(bool).
			// Check if the threshold changes with the level of KP.
			common.PrintPurpleWarning(
				fmt.Sprintf(
					"KP Bollinger band not large enough %v. No Profit anticipated. Passing",
					thresUp-thresLow,
				),
			)
			continue
		}
		var (
			long  Signal[Trade]
			short Signal[Trade]
		)
		switch {
		case premium.Premium < thresLow:
			// Enter position
			long.Type = "trade"
			short.Type = "trade"

			long.Data.Exchange = "upbit"
			long.Data.Position = "long"
			long.Data.Asset = premium.Asset

			short.Data.Exchange = "binance"
			short.Data.Position = "short"
			short.Data.Asset = premium.Asset

		case premium.Premium > thresUp:
			// Exit position
			long.Type = "trade"
			short.Type = "trade"

			long.Data.Exchange = "upbit"
			long.Data.Position = "short"
			long.Data.Asset = premium.Asset

			short.Data.Exchange = "binance"
			short.Data.Position = "long"
			short.Data.Asset = premium.Asset
		default:
			continue
		}
		compareResult = append(compareResult, long)
		compareResult = append(compareResult, short)
	}
	return compareResult, nil
}
