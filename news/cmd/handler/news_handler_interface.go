package handler

import "github.com/gin-gonic/gin"

type NewsHandlerInterface interface {
	Home(c *gin.Context)
	ShowNews(c *gin.Context)
	CreateNews(title, content, expires string) (int, []string)
}
