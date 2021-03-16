package controller

import (
	"github.com/CyganFx/snippetBox-microservice/news/pkg/domain"
	"github.com/gin-gonic/gin"
)

type NewsControllerInterface interface {
	Home(c *gin.Context)
	ShowNews(c *gin.Context)
	CreateNews(news *domain.News) (int, error)
}
