package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"snippetBox-microservice/catalog/internal/repository"
	"snippetBox-microservice/catalog/pkg/domain"
	"snippetBox-microservice/catalog/utils/helpers"
	"strconv"
)

type CatalogController struct {
	repository repository.ICatalogRepository
	helper     helpers.HelperInterface
}

func New(repository repository.ICatalogRepository, helper helpers.HelperInterface) ICatalogController {
	return &CatalogController{repository: repository, helper: helper}
}

func (c2 CatalogController) ShowProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c2.helper.NotFound(c)
		return
	}

	product, err := c2.repository.GetById(id)
	if err != nil {
		if errors.Is(err, domain.ErrNoRecord) {
			c2.helper.NotFound(c)
			return
		} else {
			c2.helper.ServerError(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, product)
}

func (c2 CatalogController) Home(c *gin.Context) {
	products, err := c2.repository.GetAll()
	if err != nil {
		c2.helper.ServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (c2 CatalogController) ShowProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	productsByCategory, err := c2.repository.GetByCategory(category)

	if err != nil {
		c2.helper.ServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, productsByCategory)
}

func (c2 CatalogController) CreateProduct(prod *domain.Product) (int, error) {
	id, errorsSlice := c2.repository.Insert(prod.Title, prod.Category, prod.Description, prod.Price)
	if errorsSlice != nil {
		return -1, errorsSlice
	}
	return id, nil
}
