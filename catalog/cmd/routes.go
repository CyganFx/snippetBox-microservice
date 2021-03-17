package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", app.catalogController.Home)
	r.GET("/catalog/:id", app.catalogController.ShowProduct)
	r.GET("/catalog/:category", app.catalogController.ShowProductsByCategory)
	r.POST("/catalog/create", app.catalogController.CreateProduct)

	return r
}
