package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"snippetBox-microservice/catalog/api/grpc/protobuffs"
	"snippetBox-microservice/catalog/cmd/grpc_client"
	"snippetBox-microservice/catalog/internal/repository"
	"snippetBox-microservice/catalog/pkg/domain"
	"snippetBox-microservice/catalog/utils/helpers"
	"strconv"
)

type CatalogController struct {
	repository repository.ICatalogRepository
	helper     helpers.HelperInterface
	grpcClient protobuffs.NewsServiceClient
}

func New(repository repository.ICatalogRepository, helper helpers.HelperInterface, grpcClient protobuffs.NewsServiceClient) ICatalogController {
	return &CatalogController{repository: repository, helper: helper, grpcClient: grpcClient}
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

func (c2 CatalogController) CreateProduct(c *gin.Context) {
	var product domain.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, "sorry, bad request")
	}

	title := product.Title
	category := product.Category
	description := product.Description
	price := product.Price

	_, err = c2.repository.Insert(title, category, description, price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "insert error")
	}
	log.Println("Grpc client is invoked")

	newsTitle := "We created a new product"
	newsContent := "Product " + title + " that is in " + category + " category " + "with price: " + fmt.Sprintf("%f", price)
	expires := "7"
	grpcClient := c2.grpcClient

	news_id := grpc_client.DoCreateNews(grpcClient, newsTitle, newsContent, expires)

	response := grpc_client.DoGetNews(grpcClient, news_id.Id)

	c.JSON(200, response)
}
