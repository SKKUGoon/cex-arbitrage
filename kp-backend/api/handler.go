package api

import (
	"context"
	"fmt"
	"kimchi/common"
	"kimchi/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Signal[T any] struct {
	Type string `json:"type" example:"iexa"`
	Data T      `json:"data"`
}

type Band struct {
	Asset string  `json:"asset" example:"BTC"`
	Upper float64 `json:"upper" example:"3.42"`
	Lower float64 `json:"lower" example:"0.56"`
}

type StatusMessage struct {
	Message string `json:"message" example:"some message"`
}

// handleBandP
// @Summary Upload normal premium boundaries
// @Description Compare duel traded crypto assets. (Traded both in upbit and binance)
// @Description Kimchi Premium will have normal - rate of premiums.
// @Description Upload the boundaries to Redis Database
// @Accept json
// @Produce json
// @Router /band [post]
// @Success 200 {object} Signal[StatusMessage]
// @Failure 400 {object} Signal[StatusMessage]
// @Failure 500 {object} Signal[StatusMessage]
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
		err1 := dao.RdbOpCreate(client, context.Background(), bandU)
		err2 := dao.RdbOpCreate(client, context.Background(), bandL)
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
