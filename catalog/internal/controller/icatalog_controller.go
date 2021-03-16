package controller

import (
	"github.com/gin-gonic/gin"
	"snippetBox-microservice/catalog/pkg/domain"
)

type ICatalogController interface {
	Home(c *gin.Context)
	ShowProduct(c *gin.Context)
	ShowProductsByCategory(c *gin.Context)
	CreateProduct(prod *domain.Product) (int, error)
}
