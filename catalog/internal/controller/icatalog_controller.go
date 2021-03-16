package controller

import (
	"github.com/gin-gonic/gin"
)

type ICatalogController interface {
	Home(c *gin.Context)
	ShowProduct(c *gin.Context)
	ShowProductsByCategory(c *gin.Context)
	CreateProduct(c *gin.Context)
}
