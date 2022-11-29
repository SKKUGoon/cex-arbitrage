package common

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleOptionalQueryParam
// args: carries query argument name and whether it is not optional
//
//	if optional
func HandleOptionalQueryParam(c *gin.Context, args map[string]bool) (map[string]string, error) {
	result := map[string]string{}
	for ask, neccesary := range args {
		i, ok := c.GetQuery(ask)
		if neccesary && !ok {
			msg := fmt.Sprintf(REQUEST_MISSING_PARAM, ask)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": msg,
			})
			return map[string]string{}, errors.New(REQUEST_MISSING_PARAM)
		}
		result[ask] = i
	}
	return result, nil
}
