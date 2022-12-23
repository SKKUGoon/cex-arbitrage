package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"kimchi/common"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

const MINIMUM_BOUND_LENGTH = 0.020

var SignalMQ SignalMessageQueue
var ctx = context.Background()

// NewSignalReciever
// Create redis pubsub client with 2 channels (SignalMessage, TradeMessage)
// @param configFile redis login info yaml file
func NewSignalReciever(configFile string) SignalMessageQueue {
	c := CacheNewConn(configFile)
	common.PrintGreenOk("Subscribe to Redis `Signal` Channel")
	return SignalMessageQueue{
		client:        c,
		SignalMessage: make(chan []byte, 1000),
		TradeMessage:  make(chan []byte, 1000),
	}
}

// getBandInfo
// Get premium bollinger band information from redis database.
// @param bandUD: whether it is upper or lower
// @param client: redis database
func getBandInfo(bandUD string, client *redis.Client) (map[string]string, error) {
	common.PrintYellowOperation("Getting band upper, lower from redis")
	// Get Band information
	var bandMap map[string]string
	var err error
	switch bandUD {
	case "upper":
		searchKeyUpper := RdbKeyField{Key: "band_upper"}
		bandMap, err = RdbOpRead(client, context.Background(), searchKeyUpper)
	case "lower":
		searchKeyLower := RdbKeyField{Key: "band_lower"}
		bandMap, err = RdbOpRead(client, context.Background(), searchKeyLower)
	}
	if err != nil {
		return nil, err
	}
	return bandMap, nil
}

// comparePremium
// If Boundary is less than `MINIMUM_BOUND_LENGTH` return empty trade signal and false
// Else based on the premium is lower than lower-bound or higher than upper-bound
// send out trading message, in Signal[Trade] format
// @param CurrentPremium: parsed incoming message.
// @param upper, lower: map[string]string that looks like {<asset name>: <boundary value>}
func comparePremium(p CurrentPremium, upper, lower map[string]string) (Position, bool) {
	common.PrintYellowOperation("Comparing currentPremium with upper and lower band information")
	var (
		pos Position
	)

	thresLow, _ := strconv.ParseFloat(lower[p.AssetPremium.Asset], 64)
	thresUp, _ := strconv.ParseFloat(upper[p.AssetPremium.Asset], 64)

	if math.Abs(thresUp-thresLow) < MINIMUM_BOUND_LENGTH {
		common.PrintPurpleWarning(
			fmt.Sprintf(
				"Asset: %s | No Profit anticipated\nBand not large enough | Size: %v",
				p.AssetPremium.Asset, thresUp-thresLow,
			),
		)
		return Position{}, false
	}

	switch {
	case p.AssetPremium.Premium < thresLow:
		// Enter position
		pos.Type = "enter"
		pos.Xlong = "upbit"
		pos.Xshort = "binance"
		pos.Asset = p.AssetPremium.Asset
		pos.PrcLong = p.AssetPremium.LongBestAskPrc
		pos.PrcShort = p.AssetPremium.ShortBestBidPrc

	case p.AssetPremium.Premium > thresUp:
		// Exit position
		pos.Type = "exit"
		pos.Xlong = "upbit"
		pos.Xshort = "binance"
		pos.Asset = p.AssetPremium.Asset
		pos.PrcLong = p.AssetPremium.LongBestAskPrc
		pos.PrcShort = p.AssetPremium.ShortBestBidPrc

	default:
		common.PrintBlueStatus(
			fmt.Sprintf(
				"Asset %s || (low) %v < %v < %v (up) || No Trade",
				p.AssetPremium.Asset,
				thresLow, p.AssetPremium.Premium, thresUp,
			),
		)
		return Position{}, false
	}
	return pos, true
}

func (mq *SignalMessageQueue) broadcastSignal(msg []byte) {
	mq.SignalMessage <- msg
}

// mqHandler *SignalMessageQueue
// infinite loop that sends out trade messages. If threshold is outside boundary
// Sendout trade messages.
// It renews its band information every minute.
func (mq *SignalMessageQueue) mqHandler() {
	// Initial Band information
	var bandUpper map[string]string
	var bandLower map[string]string
	var p Signal[CurrentPremium]
	bandUpper, err1 := getBandInfo("upper", mq.client)
	bandLower, err2 := getBandInfo("lower", mq.client)
	if err1 != nil || err2 != nil {
		log.Panicln(err1.Error() + err2.Error())
	}
	// Renewing band boundary on the count of ticker
	tic := time.NewTicker(time.Minute * 1)
	for {
		select {
		case sigM := <-mq.SignalMessage:
			// Compare signal with band
			// Send out Trade message
			_ = json.Unmarshal(sigM, &p)
			tradeSigs, ok := comparePremium(p.Data, bandUpper, bandLower)
			if ok {
				tradePacket, _ := json.Marshal(tradeSigs)
				err := mq.client.Publish(ctx, "trade_channel", tradePacket).Err()
				if err != nil {
					common.PrintPurpleWarning(err.Error())
				}
			}
			continue

		case <-tic.C:
			// Update Band information
			bandUpper, err1 = getBandInfo("upper", mq.client)
			bandLower, err2 = getBandInfo("lower", mq.client)
			if err1 != nil || err2 != nil {
				common.PrintRedError(err1.Error() + err2.Error())
			}
		}
	}
}

// Run *SignalMessageQueue
// Goroutine mqHandler. Plus insert message using `broadcastSignal` method.
func (mq *SignalMessageQueue) Run() error {
	go mq.mqHandler()
	subscriber := mq.client.Subscribe(ctx, "signal_channel")
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			log.Panicln(err)
		}
		mq.broadcastSignal([]byte(msg.Payload))
	}
}
