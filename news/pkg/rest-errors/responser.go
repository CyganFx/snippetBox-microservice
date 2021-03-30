package rest_errors

import "github.com/gin-gonic/gin"

type Responser interface {
	ServerError(c *gin.Context, err error)
	ClientError(c *gin.Context, status int)
	ClientErrorWithDescription(c *gin.Context, status int, description string)
	NotFound(c *gin.Context)
}
