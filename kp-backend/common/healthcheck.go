package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAlive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"type": "isalive",
		"data": "pong",
	})
}
