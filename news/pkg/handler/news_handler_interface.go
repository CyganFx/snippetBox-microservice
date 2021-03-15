package handler

import (
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/gin-gonic/gin"
)

type NewsHandlerInterface interface {
	Home(c *gin.Context)
	ShowNews(c *gin.Context)
	CreateNews(news *domain.News) (int, error)
}
