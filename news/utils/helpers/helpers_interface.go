package helpers

import (
	"github.com/gin-gonic/gin"
	"time"
)

type HelperInterface interface {
	HumanDate(t time.Time) string
	ServerError(c *gin.Context, err error)
	ClientError(c *gin.Context, status int)
	ClientErrorWithDescription(c *gin.Context, status int, description string)
	NotFound(c *gin.Context)
}
