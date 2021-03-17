package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", app.catalogController.Home)
	r.POST("/catalog/create", app.catalogController.CreateProduct)
	r.GET("/catalog/category/:id", app.catalogController.ShowProductsByCategory) //name of category
	r.GET("/catalog/product/:id", app.catalogController.ShowProduct)

	return r
}
