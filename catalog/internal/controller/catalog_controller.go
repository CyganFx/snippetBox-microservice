package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
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
	category := c.Param("id")
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
		c2.helper.ClientErrorWithDescription(c, http.StatusBadRequest, "type normal data please")
	}

	title := product.Title
	category := product.Category
	description := product.Description
	price := product.Price

	_, err = c2.repository.Insert(title, category, description, price)
	if err != nil {
		c2.helper.ServerError(c, err)
	}
	log.Println("Grpc client is invoked")

	newsTitle := "We created a new product"
	newsContent := "Product " + title + " that is in " + category + " category " + "with price: " + fmt.Sprintf("%f", price)
	expires := "7"
	grpcClient := c2.grpcClient

	createNewsResponse := grpc_client.DoCreateNews(grpcClient, newsTitle, newsContent, expires)

	getNewsResponse := grpc_client.DoGetNews(grpcClient, createNewsResponse.Id)

	timeTimeCreated, err := ptypes.Timestamp(getNewsResponse.Created)
	if err != nil {
		c2.helper.ServerError(c, err)
	}
	formattedCreated := c2.helper.HumanDate(timeTimeCreated)

	timeTimeExpires, err := ptypes.Timestamp(getNewsResponse.Expires)
	if err != nil {
		c2.helper.ServerError(c, err)
	}
	formattedExpires := c2.helper.HumanDate(timeTimeExpires)

	c.JSON(http.StatusOK, gin.H{
		"id":      getNewsResponse.Id,
		"title":   getNewsResponse.Title,
		"content": getNewsResponse.Content,
		"created": formattedCreated,
		"expires": formattedExpires,
	})
}
