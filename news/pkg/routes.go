package pkg

import (
	"github.com/gin-gonic/gin"
	"snippetBox-microservice/news/pkg/domain"
)

func SetupRoutes(controller domain.NewsController) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), SecureHeaders())

	r.GET("/", controller.Home)
	r.GET("/news/:id", controller.ShowNews)

	return r
}
