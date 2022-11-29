package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"kimchi/common"
	"kimchi/dao"
	"net/http"
)

// handleBandP
// @Summary premium
func handleBandP(c *gin.Context, client *redis.Client) {
	// Process signal message
	var sig Signal[[]Band]
	var resp Signal[StatusMessage]
	err := c.ShouldBindJSON(&sig)
	if err != nil {
		resp.Type = "failed"
		resp.Data.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// Band Insert in database: key's name is "band_upper" and "band_lower"
	common.PrintYellowOperation(fmt.Sprintf("Processing %v data", len(sig.Data)))
	for i := 0; i < len(sig.Data); i++ {
		bandU := dao.RdbKeyFieldValue[float64]{
			Key:   "band_upper",
			Field: sig.Data[i].Asset,
			Value: sig.Data[i].Upper,
		}
		bandL := dao.RdbKeyFieldValue[float64]{
			Key:   "band_lower",
			Field: sig.Data[i].Asset,
			Value: sig.Data[i].Lower,
		}
		err1 := dao.RdbOpCreate[float64](client, context.Background(), bandU)
		err2 := dao.RdbOpCreate[float64](client, context.Background(), bandL)
		if err1 != nil || err2 != nil {
			common.PrintRedError("Band update error")
			resp.Type = "failed"
			resp.Data.Message = err1.Error() + " " + err2.Error()
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
	}
	// Return Band Update successful
	resp.Type = "success"
	resp.Data.Message = "Band updated"
	c.JSON(http.StatusOK, resp)
}

func handlePremiumP(c *gin.Context, client *redis.Client) {
	var sig Signal[CurrentPremium]
	var resp Signal[StatusMessage]

	err := c.ShouldBindJSON(&sig)
	if err != nil {
		var resp Signal[StatusMessage]
		resp.Type = "failed"
		resp.Data.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	txOrder, err := comparePremium(sig.Data, client)
	if err != nil {
		resp.Type = "failed"
		resp.Data.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// Publish message to Redis channel
	for _, tx := range txOrder {
		txByte, err := json.Marshal(tx)
		if err != nil {
			common.PrintRedError(err.Error())
			resp.Type = "publish failed"
			resp.Data.Message = err.Error()
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		err = client.Publish(context.Background(), "trade_channel", txByte).Err()
		if err != nil {
			common.PrintRedError(err.Error())
			continue
		}
	}
	resp.Type = "success"
	resp.Data.Message = "Premium successfully evaluated"
	c.JSON(http.StatusOK, resp)
}
