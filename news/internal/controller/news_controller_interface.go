package controller

import (
	"github.com/gin-gonic/gin"
	"snippetBox-microservice/news/pkg/domain"
)

type NewsControllerInterface interface {
	Home(c *gin.Context)
	ShowNews(c *gin.Context)
	CreateNews(news *domain.News) (int, error)
}
