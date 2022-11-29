package common

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(ALLOW_ORIGIN, ALLOW_ORIGIN_VALUE)
		c.Header(ALLOW_CREDENTIALS, ALLOW_CREDENTIALS_VALUE)
		c.Next()
	}
}
