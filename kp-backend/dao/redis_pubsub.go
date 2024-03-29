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

const MINIMUM_BOUND_LENGTH = 0.04

var SignalMQ SignalMessageQueue
var ctx = context.Background()

// NewSignalReciever
// Create redis pubsub client with 2 channels (SignalMessage, TradeMessage)
// @param configFile redis login info yaml file
func NewSignalReciever(configFile string) SignalMessageQueue {
	c := CacheNewConn(configFile)
	common.PrintGreenOk("Goroutine Notice: Redis `Signal` Channel")
	common.PrintGreenOk("Goroutine Notice: Redis `Trade` Channel")
	return SignalMessageQueue{
		client:        c,
		SignalMessage: make(chan []byte, 1000),
		NoticeMessage: make(chan []byte, 1000),
		TradeMessage:  make(chan []byte, 1000),
	}
}

// getBandInfo
// Get premium bollinger band information from redis database.
// @param bandUD: whether it is upper or lower
// @param client: redis database
func getBandInfo(bandUD string, client *redis.Client) (map[string]string, error) {
	common.PrintYellowOperation("Get band upper, lower from redis")
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

// getFlaggedAsset
// Get flagged asset. Asset was flagged by being mentioned in notice board etc.
// @param client: redis database
func getFlaggedAsset(client *redis.Client) (map[string]string, error) {
	common.PrintYellowOperation("Get flagged asset")
	// Get flagged asset information
	var flagMap map[string]string
	searchKey := RdbKeyField{Key: "beware"}
	flagMap, err := RdbOpRead(client, context.Background(), searchKey)
	if err != nil {
		return nil, err
	}
	return flagMap, nil
}

// comparePremium
// If Boundary is less than `MINIMUM_BOUND_LENGTH` return empty trade signal and false
// Else based on the premium is lower than lower-bound or higher than upper-bound
// send out trading message, in Signal[Trade] format
// @param CurrentPremium: parsed incoming message.
// @param upper, lower: map[string]string that looks like {<asset name>: <boundary value>}
func comparePremium(p CurrentPremium, upper, lower, flagged map[string]string) (Position, bool) {
	// common.PrintYellowOperation("Premium Comparison")
	var (
		pos Position
	)

	thresLow, _ := strconv.ParseFloat(lower[p.AssetPremium.Asset], 64)
	thresUp, _ := strconv.ParseFloat(upper[p.AssetPremium.Asset], 64)

	switch {
	// Enter position
	case p.AssetPremium.Premium < thresLow:
		// If no profit anticipated, don't emit `enter` signal
		if math.Abs(thresUp-thresLow) < MINIMUM_BOUND_LENGTH {
			common.PrintBlueStatus(
				fmt.Sprintf(
					"| Asset: %-7v | No Profit anticipated              | BandSize: %.3f |",
					p.AssetPremium.Asset, thresUp-thresLow,
				),
			)
			return Position{}, false
		}
		if flagged[p.AssetPremium.Asset] == "no_enter" {
			common.PrintRedError(
				fmt.Sprintf(
					"| Asset: %-7v |      A S S E T  F L A G G E D      | Chk notice brd  |",
					p.AssetPremium.Asset,
				),
			)
		}
		common.PrintGreenOk(
			fmt.Sprintf(
				"| Asset: %-7v | Lower than thres %.4f > %.4f   | BandSize: %.3f |",
				p.AssetPremium.Asset, thresLow, p.AssetPremium.Premium, thresUp-thresLow,
			),
		)
		pos.Type = "enter"
		pos.Xlong = "upbit"
		pos.Xshort = "binance"
		pos.Asset = p.AssetPremium.Asset
		pos.PrcLong = p.AssetPremium.LongBestAskPrc
		pos.PrcShort = p.AssetPremium.ShortBestBidPrc
		pos.RptPremium = p.AssetPremium.Premium

	// Exit position
	case p.AssetPremium.Premium > (thresUp * 3 / 5):
		// Exit signal should be made regardless of band size.
		// since the band size might have been reduced after the
		// opening of the position.
		common.PrintGreenOk(
			fmt.Sprintf(
				"| Asset: %-7v | Higher than thres %.4f < %.4f  | BandSize: %.3f |",
				p.AssetPremium.Asset, thresUp, p.AssetPremium.Premium, thresUp-thresLow,
			),
		)
		pos.Type = "exit"
		pos.Xlong = "upbit"
		pos.Xshort = "binance"
		pos.Asset = p.AssetPremium.Asset
		pos.PrcLong = p.AssetPremium.LongBestAskPrc
		pos.PrcShort = p.AssetPremium.ShortBestBidPrc
		pos.RptPremium = p.AssetPremium.Premium

	// Between the band - No position
	default:
		bInfo := fmt.Sprintf(
			"(low) %.3f < %.3f < %.3f (up)",
			thresLow, p.AssetPremium.Premium, thresUp,
		)
		common.PrintBlueStatus(
			fmt.Sprintf(
				"| Asset: %-7v | %-34v | Middle No Trade |",
				p.AssetPremium.Asset,
				bInfo,
			),
		)
		return Position{}, false
	}
	return pos, true
}

func (mq *SignalMessageQueue) broadcastSignal(msg []byte) {
	mq.SignalMessage <- msg
}

func (mq *SignalMessageQueue) broadcastNotice(msg []byte) {
	mq.NoticeMessage <- msg
}

// mqHandler *SignalMessageQueue
// infinite loop that sends out trade messages. If threshold is outside boundary
// Sendout trade messages.
// It renews its band information every minute.
func (mq *SignalMessageQueue) mqHandler() {
	// Initial Band information
	var bandUpper map[string]string
	var bandLower map[string]string
	var flaggedAsset map[string]string
	bandUpper, err1 := getBandInfo("upper", mq.client)
	bandLower, err2 := getBandInfo("lower", mq.client)
	flaggedAsset, err3 := getFlaggedAsset(mq.client)
	if err1 != nil || err2 != nil || err3 != nil {
		log.Panicln("1", err1.Error(), "2", err2.Error(), "3", err3.Error())
	}
	// Renewing band boundary on the count of ticker
	tic := time.NewTicker(time.Minute * 1)
	for {
		select {
		case sigM := <-mq.SignalMessage:
			// Compare signal with band
			var p Signal[CurrentPremium]
			err := json.Unmarshal(sigM, &p)
			if err != nil {
				common.PrintPurpleWarning("malfored signal_channel message")
				continue
			}
			tradeSigs, ok := comparePremium(p.Data, bandUpper, bandLower, flaggedAsset)
			if ok {
				// Send out Trade message
				tradePacket, _ := json.Marshal(tradeSigs)
				err := mq.client.Publish(ctx, "trade_channel", tradePacket).Err()
				if err != nil {
					common.PrintPurpleWarning(err.Error())
				}
			}
			continue

		case notM := <-mq.NoticeMessage:
			fmt.Println(flaggedAsset)
			// Notice Board trading message
			var p Signal[BlockNotice]
			var tradeSigs Position
			err := json.Unmarshal(notM, &p)
			if err != nil {
				common.PrintPurpleWarning("malformed notice_channel message")
				continue
			}
			switch {
			case !p.Data.Complete:
				// Add to redis database
				bewareAsset := RdbKeyFieldValue[string]{
					Key:   "beware",
					Field: p.Data.Asset,
					Value: "no_enter",
				}
				err := RdbOpCreate(mq.client, ctx, bewareAsset)
				if err != nil {
					common.PrintPurpleWarning("fail to ADD beware asset" + p.Data.Asset)
				}
				// Enter position. Upbit closing their transfer
				tradeSigs = Position{
					Type:     "enter",
					Xlong:    "upbit",
					Xshort:   "binance",
					Asset:    p.Data.Asset,
					PrcLong:  -1,
					PrcShort: -1,
				}
				flaggedAsset, err3 = getFlaggedAsset(mq.client)
				if err3 != nil {
					common.PrintRedError(err3.Error())
				}
			case p.Data.Complete:
				// Delete from redis database
				bewareAssetDel := RdbKeyField{
					Key:   "beware",
					Field: p.Data.Asset,
				}
				err := RdbOpDelete(mq.client, ctx, bewareAssetDel)
				if err != nil {
					common.PrintPurpleWarning("fail to REMOVE beware asset" + p.Data.Asset)
				}
				// Exit position. Upbit finished their maintenance
				tradeSigs = Position{
					Type:     "exit",
					Xlong:    "upbit",
					Xshort:   "binance",
					Asset:    p.Data.Asset,
					PrcLong:  -1,
					PrcShort: -1,
				}
				flaggedAsset, err3 = getFlaggedAsset(mq.client)
				if err3 != nil {
					common.PrintRedError(err3.Error())
				}
			}
			// Send out Trade message
			tradePacket, _ := json.Marshal(tradeSigs)
			err = mq.client.Publish(ctx, "trade_channel", tradePacket).Err()
			if err != nil {
				common.PrintPurpleWarning(err.Error())
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

func (mq *SignalMessageQueue) mqToSigChan() {
	subscriber := mq.client.Subscribe(ctx, "signal_channel")
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			common.PrintRedError("redis pubsub on signal_channel receive error", err.Error())
		}
		mq.broadcastSignal([]byte(msg.Payload))
	}
}

func (mq *SignalMessageQueue) mqToNotChan() {
	subscriber := mq.client.Subscribe(ctx, "notice_channel")
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			common.PrintRedError("redis pubsub on notice_channel receive error", err.Error())
		}
		mq.broadcastNotice([]byte(msg.Payload))
	}
}

// Run *SignalMessageQueue
// Goroutine mqHandler. Plus insert message using `broadcastSignal` method.
func (mq *SignalMessageQueue) Run() {
	// c1, cancel := context.WithCancel(context.Background())
	go mq.mqToSigChan()
	go mq.mqToNotChan()
	mq.mqHandler()
}
